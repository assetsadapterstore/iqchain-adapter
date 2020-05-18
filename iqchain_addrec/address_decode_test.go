package iqchain_addrec

import (
	"fmt"
	"github.com/ArkEcosystem/go-crypto/crypto"
	"testing"
)


func TestAdmSymbolList(t * testing.T){
	address, _ := crypto.AddressFromPassphrase("")
	fmt.Println(address)
}