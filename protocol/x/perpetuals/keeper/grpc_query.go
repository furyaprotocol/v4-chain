package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/perpetuals/types"
)

var _ types.QueryServer = Keeper{}
