package iqchain_addrec

import (
	"fmt"
	"github.com/assetsadapterstore/iqchain-adapter/sdk/crypto"
	"testing"
)


func TestAdmSymbolList(t * testing.T){
	address, _ := crypto.ValidateAddress("QSBVzepRruTCy3ezd1XHvXjmg8yYbU65eC")
	fmt.Println(address)
}