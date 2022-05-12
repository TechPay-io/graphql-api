// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"math/big"
	"techpay-api-graphql/internal/repository"
	"techpay-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TpcBurnedTotal resolves total amount of burned TPC tokens in WEI units.
func (rs *rootResolver) TpcBurnedTotal() hexutil.Big {
	val, err := repository.R().TpcBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return hexutil.Big{}
	}
	return hexutil.Big(*new(big.Int).Mul(big.NewInt(val), types.BurnDecimalsCorrection))
}

// TpcBurnedTotalAmount resolves total amount of burned TPC tokens in TPC units.
func (rs *rootResolver) TpcBurnedTotalAmount() float64 {
	val, err := repository.R().TpcBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return 0
	}
	return float64(val) / types.BurnTPCDecimalsCorrection
}

// TpcLatestBlockBurnList resolves a list of the latest block burns.
func (rs *rootResolver) TpcLatestBlockBurnList(args struct{ Count int32 }) ([]types.TpcBurn, error) {
	if args.Count < 1 || args.Count > 50 {
		args.Count = 25
	}
	return repository.R().TpcBurnList(int64(args.Count))
}
