package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/sending/types"
)

var _ types.QueryServer = Keeper{}
