package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
