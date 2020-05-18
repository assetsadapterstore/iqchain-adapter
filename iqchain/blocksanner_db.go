package iqchain

import (
	"fmt"
	"github.com/assetsadapterstore/iqchain-adapter/sdk/client"
	"github.com/blocktree/openwallet/v2/openwallet"
)

//SaveLocalBlockHead 记录区块高度和hash到本地
func (bs *IQCBlockScanner) SaveLocalBlockHead(blockHeight uint64, blockHash string) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	header := &openwallet.BlockHeader{
		Hash:   blockHash,
		Height: uint64(blockHeight),
		Fork:   false,
		Symbol: bs.wm.Symbol(),
	}

	return bs.BlockchainDAI.SaveCurrentBlockHead(header)
}

//GetLocalBlockHead 获取本地记录的区块高度和hash
func (bs *IQCBlockScanner) GetLocalBlockHead() (uint64, string, error) {

	if bs.BlockchainDAI == nil {
		return 0, "", fmt.Errorf("Blockchain DAI is not setup ")
	}

	header, err := bs.BlockchainDAI.GetCurrentBlockHead(bs.wm.Symbol())
	if err != nil {
		return 0, "", err
	}

	return uint64(header.Height), header.Hash, nil
}

//SaveLocalBlock 记录本地新区块
func (bs *IQCBlockScanner) SaveLocalBlock(blockHeader *client.Block) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	header := &openwallet.BlockHeader{
		Hash:              blockHeader.Id,
		Previousblockhash: blockHeader.Previous,
		Height:            uint64(blockHeader.Height),
		Time:              uint64(blockHeader.Timestamp.Unix),
		Symbol:            bs.wm.Symbol(),
	}

	return bs.BlockchainDAI.SaveLocalBlockHead(header)
}

//GetLocalBlock 获取本地区块数据
func (bs *IQCBlockScanner) GetLocalBlock(height uint64) (*client.Block, error) {

	if bs.BlockchainDAI == nil {
		return nil, fmt.Errorf("Blockchain DAI is not setup ")
	}

	header, err := bs.BlockchainDAI.GetLocalBlockHeadByHeight(uint64(height), bs.wm.Symbol())
	if err != nil {
		return nil, err
	}

	block := &client.Block{
		Id:     header.Hash,
		Height: int64(header.Height),
	}

	return block, nil
}

//SaveUnscanRecord 保存交易记录到钱包数据库
func (bs *IQCBlockScanner) SaveUnscanRecord(record *openwallet.UnscanRecord) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.SaveUnscanRecord(record)
}

//DeleteUnscanRecord 删除指定高度的未扫记录
func (bs *IQCBlockScanner) DeleteUnscanRecord(height uint64) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.DeleteUnscanRecordByHeight(uint64(height), bs.wm.Symbol())
}

func (bs *IQCBlockScanner) GetUnscanRecords() ([]*openwallet.UnscanRecord, error) {

	if bs.BlockchainDAI == nil {
		return nil, fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.GetUnscanRecords(bs.wm.Symbol())
}
