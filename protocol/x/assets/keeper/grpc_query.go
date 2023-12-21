package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/assets/types"
)

var _ types.QueryServer = Keeper{}
