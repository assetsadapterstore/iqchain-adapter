package iqchain

import (
	"fmt"
	"github.com/blocktree/openwallet/v2/common"
	"github.com/blocktree/openwallet/v2/crypto"
	"github.com/blocktree/openwallet/v2/openwallet"
)

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	BlockID     string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *openwallet.UnscanRecord {
	obj := openwallet.UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}
