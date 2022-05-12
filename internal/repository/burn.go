/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Photon/Sirius full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"techpay-api-graphql/internal/types"
)

// StoreTpcBurn stores the given native TPC burn per block record into the persistent storage.
func (p *proxy) StoreTpcBurn(burn *types.TpcBurn) error {
	p.cache.TpcBurnUpdate(burn, p.db.BurnTotal)
	return p.db.StoreBurn(burn)
}

// TpcBurnTotal provides the total amount of burned native TPC.
func (p *proxy) TpcBurnTotal() (int64, error) {
	return p.cache.TpcBurnTotal(p.db.BurnTotal)
}

// TpcBurnList provides list of per-block burned native TPC tokens.
func (p *proxy) TpcBurnList(count int64) ([]types.TpcBurn, error) {
	return p.db.BurnList(count)
}
