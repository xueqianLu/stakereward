package models

import "math/big"

type LockInfo struct {
	UserAddress string   `json:"useraddress"`
	LockAmount  *big.Int `json:"lockamount"`
	LockBlock   uint64   `json:"lockblock"`
	LockTime    uint64   `json:"locktime"`
}

type WithDrawInfo struct {
	WithdrawBlock uint64 `json:"withdrawblock"`
	WithdrawTime  uint64 `json:"withdrawtime"`
}

type RecordDetail struct {
	*LockInfo
	*WithDrawInfo
	InDeposit bool `json:"instaking"`
}

type NodeInfo struct {
	NodeAddress string         `json:"nodeaddress"`
	Details     []RecordDetail `json:"details"`
}
