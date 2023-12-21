package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/vest/types"
)

var _ types.QueryServer = Keeper{}
