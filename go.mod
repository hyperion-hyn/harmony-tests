module github.com/hyperion/hyperion-tests

go 1.14

replace github.com/hyperion-hyn/hyperion-tf => ../hyperion-tf
replace github.com/ethereum/go-ethereum => ../go-ethereum

require (
	github.com/cosmos/cosmos-sdk v0.39.0 // indirect
	github.com/ethereum/go-ethereum v1.8.27
	github.com/hyperion-hyn/hyperion-tf v0.0.0-00010101000000-000000000000
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
)
