package async

import (
	"context"
	"fmt"
	"githu.com/stakereward/contracts/hpblock"
	"githu.com/stakereward/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

var (
	bigOne     = big.NewInt(1)
	bigTen     = big.NewInt(10)
	bigHundred = big.NewInt(100)
	bigK       = big.NewInt(1000)
	big5K      = big.NewInt(5000)

	secondOneYear    = big.NewInt(60 * 60 * 24 * 365)
	rewardOneYear, _ = new(big.Int).SetString("3000000000000000000000", 10) // 30000 * 0.1 * 10^18
	rewardPerSecond  = new(big.Int).Div(rewardOneYear, secondOneYear)
)

type logHandler func(log types.Log, ctx *PullEvent) error

type PullEvent struct {
	ctx             context.Context
	client          *ethclient.Client
	startBlock      *big.Int
	endBlock        *big.Int
	lastBlock       *big.Int
	endTime         uint64
	contractList    []common.Address
	contractHandler map[common.Address]logHandler
	lockContract    *hpblock.Hpblock
	nodeInfos       map[string]*models.NodeInfo
}

func (this *PullEvent) GetTimeById(number *big.Int) (uint64, error) {
	block, err := this.client.BlockByNumber(this.ctx, number)
	if err != nil {
		return 0, err
	}
	return block.Time(), nil
}

func (this *PullEvent) AddLockInfo(nodeaddr string, info *models.LockInfo) {
	nodeinfo, exist := this.nodeInfos[nodeaddr]
	if !exist {
		nodeinfo = &models.NodeInfo{Details: make([]models.RecordDetail, 0), NodeAddress: nodeaddr}
	}
	nodeinfo.Details = append(nodeinfo.Details, models.RecordDetail{
		LockInfo:     info,
		WithDrawInfo: nil,
		InDeposit:    true,
	})
	this.nodeInfos[nodeaddr] = nodeinfo
}

func (this *PullEvent) AddWithDrawInfo(nodeaddr string, info *models.WithDrawInfo) {
	nodeinfo, exist := this.nodeInfos[nodeaddr]
	if !exist {
		logs.Error("got withdraw event but have no lock event before.")
		return
	} else {
		idx := len(nodeinfo.Details) - 1
		if nodeinfo.Details[idx].WithDrawInfo != nil {
			logs.Error("got withdraw event but there have a withdraw info exist in latest detail")
		} else {
			nodeinfo.Details[idx].WithDrawInfo = info
			nodeinfo.Details[idx].InDeposit = false
		}
	}
}

func getBlockNumer(name string) *big.Int {
	number := beego.AppConfig.String(name)
	blockNumber, _ := new(big.Int).SetString(number, 10)
	return blockNumber
}

func NewTask() *PullEvent {
	var err error
	rpc := beego.AppConfig.String("url")
	if len(rpc) == 0 {
		rpc = "https://hpbnode.com"
	}
	pullTask := &PullEvent{
		contractHandler: make(map[common.Address]logHandler),
		nodeInfos:       make(map[string]*models.NodeInfo),
	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(fmt.Sprintf("ethClient dial failed, err:", err))
	} else {
		pullTask.client = client
		pullTask.startBlock = getBlockNumer("startblock")
		pullTask.endBlock = getBlockNumer("endblock")
		pullTask.lastBlock = pullTask.startBlock
	}
	pullTask.ctx = context.Background()
	pullTask.contractList = make([]common.Address, 0)
	{
		addrs := beego.AppConfig.String("lockcontract")
		{
			if len(addrs) > 0 {
				addr := common.HexToAddress(addrs)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HpbLockHandler
				pullTask.lockContract, _ = hpblock.NewHpblock(addr, client)
				pullTask.client = client
				logs.Info("add contract to fillter log", "address ", addrs)
			}
		}
	}
	return pullTask
}

func (p *PullEvent) GetDetails() map[string]*models.NodeInfo {
	return p.nodeInfos
}

func (p *PullEvent) GetRewardInfo(nodeinfo *models.NodeInfo) []*models.RewardInfo {
	var err error
	if p.endTime == 0 {
		p.endTime, err = p.GetTimeById(p.endBlock)
		if err != nil {
			logs.Error("get endtime failed ", "err", err)
			return nil
		}
	}
	var reward = make([]*models.RewardInfo, 0)
	for i := 0; i < len(nodeinfo.Details); i++ {
		detail := nodeinfo.Details[i]
		r := &models.RewardInfo{}
		r.RecordDetail = &nodeinfo.Details[i]
		r.UserAddress = r.RecordDetail.UserAddress
		if detail.WithDrawInfo == nil {
			r.StakeDuration = p.endTime - r.RecordDetail.LockTime
		} else {
			r.StakeDuration = r.RecordDetail.WithdrawTime - r.RecordDetail.LockTime
		}
		r.Amount = new(big.Int).Mul(rewardPerSecond, big.NewInt(int64(r.StakeDuration)))
		reward = append(reward, r)
	}
	return reward
}

func (p *PullEvent) SyncLogs() {
	query := ethereum.FilterQuery{}
	query.FromBlock = p.lastBlock
	query.ToBlock = new(big.Int).Add(p.lastBlock, big.NewInt(1))
	query.Addresses = p.contractList
	for {
		var finish = false
		query.FromBlock = p.lastBlock

		//logs.Info("start fileter start at ", p.lastBlock.Text(10))
		height, err := p.client.BlockNumber(p.ctx)
		if err != nil {
			logs.Error("get block number failed", "err", err)
		}
		if height <= p.lastBlock.Uint64() {
			time.Sleep(time.Second)
			continue
		} else if (height - 5000) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, big5K)
		} else if (height - 1000) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigK)
		} else if (height - 100) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigHundred)
		} else if (height - 10) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigTen)
		} else {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigOne)
		}
		if query.ToBlock.Cmp(p.endBlock) > 0 {
			query.ToBlock = p.endBlock
			finish = true
		}

		allLogs, err := p.client.FilterLogs(p.ctx, query)
		if err != nil {
			logs.Error("filter logs failed", err)
			continue
		}
		if len(allLogs) > 0 {
			for _, vlog := range allLogs {
				handle, exist := p.contractHandler[vlog.Address]
				if exist {
					handle(vlog, p)
				}
			}
		}
		p.lastBlock = new(big.Int).Add(query.ToBlock, bigOne)
		if finish {
			logs.Info("get logs finished")
			return
		}
	}
}
