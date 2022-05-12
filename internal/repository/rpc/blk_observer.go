/*
Package rpc implements bridge to Sirius full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Photon/Sirius node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Sirius RPC interface for remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Sirius RPC interface with connection limited to specified endpoints.

We strongly discourage opening Sirius RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum"
)

// tpcHeadsObserverSubscribeTick represents the time between subscription attempts.
const tpcHeadsObserverSubscribeTick = 30 * time.Second

// observeBlocks collects new blocks from the blockchain network
// and posts them into the proxy channel for processing.
func (tpc *TpcBridge) observeBlocks() {
	var sub ethereum.Subscription
	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
		tpc.log.Noticef("block observer done")
		tpc.wg.Done()
	}()

	sub = tpc.blockSubscription()
	for {
		// re-subscribe if the subscription ref is not valid
		if sub == nil {
			tm := time.NewTimer(tpcHeadsObserverSubscribeTick)
			select {
			case <-tpc.sigClose:
				return
			case <-tm.C:
				sub = tpc.blockSubscription()
				continue
			}
		}

		// use the subscriptions
		select {
		case <-tpc.sigClose:
			return
		case err := <-sub.Err():
			tpc.log.Errorf("block subscription failed; %s", err.Error())
			sub = nil
		}
	}
}

// blockSubscription provides a subscription for new blocks received
// by the connected blockchain node.
func (tpc *TpcBridge) blockSubscription() ethereum.Subscription {
	sub, err := tpc.rpc.EthSubscribe(context.Background(), tpc.headers, "newHeads")
	if err != nil {
		tpc.log.Criticalf("can not observe new blocks; %s", err.Error())
		return nil
	}
	return sub
}
