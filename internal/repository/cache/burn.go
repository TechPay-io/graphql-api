// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fmt"
	"sync/atomic"
	"techpay-api-graphql/internal/types"
)

// burnedTotalReloadBlockPeriod represents the number of blocks pass before burned total is refreshed from the database.
const burnedTotalReloadBlockPeriod = 1200

// burnTotalContainer represents a container for collecting total burned amount used to speed up resolving the total.
type burnTotalContainer struct {
	block    uint64
	value    int64
	nextLoad uint64
}

// burnContainer is the in-memory container for burned total amount.
var burnContainer *burnTotalContainer

// TpcBurnTotal returns the current amount of total burned TPC.
func (b *MemBridge) TpcBurnTotal(loader func() (int64, error)) (int64, error) {
	// we may not have the reference yet
	if burnContainer == nil {
		return loader()
	}

	v := atomic.LoadInt64(&burnContainer.value)
	return v, nil
}

// TpcBurnUpdate updates in-memory value of the burned TPCs.
func (b *MemBridge) TpcBurnUpdate(burn *types.TpcBurn, loader func() (int64, error)) {
	// make sure we have the container properly loaded and fresh
	if burnContainer == nil || (burnContainer != nil && burnContainer.nextLoad <= uint64(burn.BlockNumber)) {
		err := b.refreshBurnUpdate(burn, loader)
		if err != nil {
			b.log.Criticalf("total burned TPC value not available; %s", err.Error())
			return
		}
	}

	// new burn received?
	if uint64(burn.BlockNumber) <= burnContainer.block {
		return
	}

	// update the value we keep
	burnContainer.block = uint64(burn.BlockNumber)
	v := atomic.LoadInt64(&burnContainer.value)
	atomic.StoreInt64(&burnContainer.value, v+burn.Value())
}

// TpcBurnUpdate updates in-memory value of the burned TPCs.
func (b *MemBridge) refreshBurnUpdate(burn *types.TpcBurn, loader func() (int64, error)) error {
	if burn.BlockNumber == 0 {
		return fmt.Errorf("zero block can not be used to load burns")
	}

	// try to load existing value
	v, err := loader()
	if err != nil {
		return err
	}

	burnContainer = &burnTotalContainer{
		block:    uint64(burn.BlockNumber) - 1,
		nextLoad: uint64(burn.BlockNumber) + burnedTotalReloadBlockPeriod,
		value:    v,
	}
	return nil
}
