package keeper

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/delaymsg/types"
)

var _ types.QueryServer = Keeper{}
