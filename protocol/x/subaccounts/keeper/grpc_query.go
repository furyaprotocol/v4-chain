package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/subaccounts/types"
)

var _ types.QueryServer = Keeper{}
