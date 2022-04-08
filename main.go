package main

import (
	"encoding/json"
	"githu.com/stakereward/async"
	"githu.com/stakereward/db"
	"github.com/astaxie/beego/logs"
	"math/big"
)

func main() {
	err := db.Init()
	if err != nil {
		logs.Error("open database failed, err ", err)
		return
	}
	db.CreateRewardInfoTable()
	task := async.NewTask()
	task.SyncLogs()
	nodeinfos := task.GetDetails()
	amount := big.NewInt(0)
	for _, info := range nodeinfos {

		reward := task.GetRewardInfo(info)
		d, _ := json.Marshal(reward)
		logs.Info("reward for ", info.NodeAddress, "detail", string(d))
		for _, r := range reward {
			var data = &db.RewardInfo{}
			data.FromDetail(info.NodeAddress, r)
			err := data.Create()
			if err != nil {
				logs.Error("insert data failed", "err", err)
			} else {
				logs.Info("insert data succeed")
			}
			amount = new(big.Int).Add(amount, r.Amount)
		}
	}
	logs.Info("total reward ", amount.Text(10))
}
