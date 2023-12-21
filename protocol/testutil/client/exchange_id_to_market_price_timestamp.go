package client

import (
	"github.com/furyanprotocol/v4-chain/protocol/daemons/pricefeed/client/types"
)

// ExchangeIdMarketPriceTimestamp contains an `ExchangeId` and an associated
// `types.MarketPriceTimestamp`. This type exists for convenience and clarity in testing the
// pricefeed client.
type ExchangeIdMarketPriceTimestamp struct {
	ExchangeId           types.ExchangeId
	MarketPriceTimestamp *types.MarketPriceTimestamp
}
