module github.com/assetsadapterstore/iqchain-adapter

go 1.12

require (
	github.com/ArkEcosystem/go-crypto v0.0.0-20200210042049-1901e1f9967a
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/Sereal/Sereal v0.0.0-20200417095951-15946cd26aa7 // indirect
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owcdrivers v1.2.1
	github.com/blocktree/go-owcrypt v1.1.4
	github.com/blocktree/openwallet/v2 v2.0.2
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422
	github.com/fatih/structs v1.1.0
	github.com/google/go-querystring v1.0.0
	github.com/hbakhtiyor/schnorr v0.1.0
	github.com/imroc/req v0.2.4
	github.com/pkg/errors v0.8.1
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/gjson v1.3.5
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876
)

//replace github.com/blocktree/openwallet => ../../openwallet
