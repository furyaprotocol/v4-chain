package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/prices/types"
)

var _ types.QueryServer = Keeper{}
