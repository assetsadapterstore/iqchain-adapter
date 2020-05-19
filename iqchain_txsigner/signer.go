package iqchain_txsigner

import (
	"github.com/assetsadapterstore/iqchain-adapter/sdk/crypto"
)

var Default = &TransactionSigner{}

type TransactionSigner struct {
}

// SignTransactionHash 交易哈希签名算法
// required
func (singer *TransactionSigner) SignTransactionHash(msg []byte, privateKey []byte, eccType uint32) ([]byte, error) {
	//_, err := owcrypt.Signature(privateKey, nil, 0, msg, uint16(len(msg)), eccType)
	//if err != owcrypt.SUCCESS {
	//	return nil, fmt.Errorf("ECC sign hash failed")
	//}
	pk := crypto.PrivateKeyFromBytes(privateKey)
	signature2, _ := pk.Sign(msg)
	return signature2, nil
}





func (singer *TransactionSigner) SignSerialize(sig []byte)[]byte {

	sb := sig[32:]
	rb := sig[:32]
	length := 6 + len(rb) + len(sb)
	b := make([]byte, length)
	b[0] = 0x30
	b[1] = byte(length - 2)
	b[2] = 0x02
	b[3] = byte(len(rb))
	offset := copy(b[4:], rb) + 4
	b[offset] = 0x02
	b[offset+1] = byte(len(sb))
	copy(b[offset+2:], sb)

	return b
}



