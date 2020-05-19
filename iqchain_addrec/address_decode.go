package iqchain_addrec

import (
	"github.com/assetsadapterstore/iqchain-adapter/sdk/crypto"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
}

// GetAddressFromPublicKey takes a Lisk public key and returns the associated address
func GetAddressFromPublicKey(publicKey []byte) string {
	//publicKeyHash := sha256.Sum256(publicKey)

	pk, err := crypto.PublicKeyFromBytes(publicKey)
	if err != nil {
		log.Error(err)
	}
	return pk.ToAddress()

}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {
	address := GetAddressFromPublicKey(hash)
	return address, nil
}



// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	veri,err := crypto.ValidateAddress(address)
	if err == nil && veri {
		return true
	}
	return false
}
