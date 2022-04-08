package models

import "math/big"

type RewardInfo struct {
	*RecordDetail
	StakeDuration uint64   `json:"stakeduration"`
	Amount        *big.Int `json:"amount"`
	RewardTx      string   `json:"rewardtx"`
}
