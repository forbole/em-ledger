module github.com/e-money/em-ledger

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/emirpasic/gods v1.12.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.4
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega v1.4.3
	github.com/regen-network/cosmos-proto v0.1.1-0.20200213154359-02baa11ea7c2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.3
	github.com/tendermint/tm-db v0.5.0
	github.com/tidwall/gjson v1.3.2
	github.com/tidwall/sjson v1.0.4
)

// replace github.com/cosmos/cosmos-sdk => ./tmpvendor/cosmos-sdk

// replace github.com/tendermint/tendermint => ./tmpvendor/tendermint
