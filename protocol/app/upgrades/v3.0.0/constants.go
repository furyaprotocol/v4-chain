package v_3_0_0

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/furyanprotocol/v4-chain/protocol/app/upgrades"
	ratelimittypes "github.com/furyanprotocol/v4-chain/protocol/x/ratelimit/types"
)

const (
	UpgradeName = "v3.0.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			ratelimittypes.StoreKey,
		},
	},
}
