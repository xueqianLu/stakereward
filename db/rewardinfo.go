package db

import (
	"githu.com/stakereward/common"
	"githu.com/stakereward/models"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RewardInfo struct {
	gorm.Model
	NodeAddr      string  `gorm:"column:nodeaddr" json:"nodeaddr"`
	UserAddr      string  `gorm:"column:useraddr" json:"useraddr"`
	LockAmount    float64 `gorm:"column:lockamount" json:"lockamount"`
	LockBlock     uint64  `gorm:"column:lockblock" json:"lockblock"`
	LockTime      uint64  `gorm:"column:locktime" json:"locktime"`
	WithdrawBlock uint64  `gorm:"column:withdrawblock" json:"withdrawblock"`
	WithdrawTime  uint64  `gorm:"column:withdrawtime" json:"withdrawtime"`
	InDeposit     bool    `gorm:"column:instaking" json:"instaking"`
	StakeDuration uint64  `json:"stakeduration"`
	Amount        float64 `gorm:"column:amount" json:"amount"`
	RewardTx      string  `gorm:"column:rewardtx" json:"rewardtx"`
}

func CreateRewardInfoTable() {
	gorm := GetORM()
	if err := gorm.AutoMigrate(&RewardInfo{}); err != nil {
		logs.Error("create table failed", "err", err)
	}
	{
		logs.Info("table exist ?", gorm.Migrator().HasTable(&RewardInfo{}))
	}
}

func (dt *RewardInfo) TableName() string {
	return "tb_rewardinfo"
}

func (dt *RewardInfo) Create() error {
	orm := GetORM()

	if err := orm.Model(&RewardInfo{}).Create(dt).Error; err != nil {
		logs.Error("db create RewardInfo error", "err", err)
		return errors.Wrap(err, "db create RewardInfo error")
	}

	return nil
}

func (dt *RewardInfo) GetAllInfo() (all []*RewardInfo, err error) {
	orm := GetORM()

	if err = orm.Model(&RewardInfo{}).Find(&all).Error; err != nil {
		logs.Error("GetAllInfo error", err.Error())
		return all, errors.Wrap(err, "GetAllInfo error")
	}

	return all, nil
}

func (dt *RewardInfo) FromDetail(nodeaddr string, detail *models.RewardInfo) {
	if dt == nil {
		panic("object is nil")
	}
	dt.NodeAddr = nodeaddr
	dt.UserAddr = detail.UserAddress
	dt.LockAmount = common.BigAmountFormatToFloat64OfDecimal(*detail.LockAmount, 18)
	dt.LockBlock = detail.LockBlock
	dt.LockTime = detail.LockTime
	if detail.WithDrawInfo != nil {
		dt.WithdrawBlock = detail.WithdrawBlock
		dt.WithdrawTime = detail.WithdrawTime
	}
	dt.InDeposit = detail.InDeposit
	dt.StakeDuration = detail.StakeDuration
	dt.Amount = common.BigAmountFormatToFloat64OfDecimal(*detail.Amount, 18)
	dt.RewardTx = detail.RewardTx
}
