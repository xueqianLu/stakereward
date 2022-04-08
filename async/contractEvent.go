package async

import (
	"githu.com/stakereward/models"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

var (
	LockEvent     = "0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34"
	WithDrawEvent = "0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4"
)

func HpbLockHandler(vLog types.Log, ctx *PullEvent) error {
	method := vLog.Topics[0]
	switch method.String() {
	case LockEvent:
		info, err := ctx.lockContract.ParseLockBal(vLog)
		if err != nil {
			logs.Error("parse lockbal log failed", "err", err)
			return err
		}
		logs.Info("get lock bal event", "user", info.Addr.String(), "node", info.NodeAddr.String())
		lockInfo := &models.LockInfo{
			UserAddress: info.Addr.String(),
			LockAmount:  new(big.Int).Set(info.LockNum),
			LockBlock:   info.Raw.BlockNumber,
		}
		number := big.NewInt(int64(info.Raw.BlockNumber))
		logs.Info("current number is ", number)
		tm, err := ctx.GetTimeById(number)
		if err != nil {
			logs.Error("get time by block number failed", "err", err)
		} else {
			lockInfo.LockTime = tm
		}
		ctx.AddLockInfo(strings.ToLower(info.NodeAddr.String()), lockInfo)
	case WithDrawEvent:
		info, err := ctx.lockContract.ParseWithdrawBal(vLog)
		if err != nil {
			logs.Error("parse withdraw bal log failed", "err", err)
			return err
		}
		logs.Info("get withdraw bal event", "user", info.Addr.String(), "node", info.NodeAddr.String())
		withDrawInfo := &models.WithDrawInfo{
			WithdrawBlock: info.Raw.BlockNumber,
		}
		tm, err := ctx.GetTimeById(big.NewInt(int64(info.Raw.BlockNumber)))
		if err != nil {
			logs.Error("get time by block number failed", "err", err)
		} else {
			withDrawInfo.WithdrawTime = tm
		}
		ctx.AddWithDrawInfo(strings.ToLower(info.NodeAddr.String()), withDrawInfo)
	}
	return nil
}
