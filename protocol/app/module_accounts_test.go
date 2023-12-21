package app_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/furyanprotocol/v4-chain/protocol/app"
	bridgemoduletypes "github.com/furyanprotocol/v4-chain/protocol/x/bridge/types"
	clobmoduletypes "github.com/furyanprotocol/v4-chain/protocol/x/clob/types"
	delaymsgtypes "github.com/furyanprotocol/v4-chain/protocol/x/delaymsg/types"
	rewardsmoduletypes "github.com/furyanprotocol/v4-chain/protocol/x/rewards/types"
	satypes "github.com/furyanprotocol/v4-chain/protocol/x/subaccounts/types"
	vestmoduletypes "github.com/furyanprotocol/v4-chain/protocol/x/vest/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAccountsToAddresses(t *testing.T) {
	expectedModuleAccToAddresses := map[string]string{
		authtypes.FeeCollectorName:                   "furya17xpfvakm2amg962yls6f84z3kell8c5leqdyt2",
		bridgemoduletypes.ModuleName:                 "furya1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv",
		distrtypes.ModuleName:                        "furya1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg",
		stakingtypes.BondedPoolName:                  "furya1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq",
		stakingtypes.NotBondedPoolName:               "furya1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605",
		govtypes.ModuleName:                          "furya10d07y265gmmuvt4z0w9aw880jnsr700jnmapky",
		ibctransfertypes.ModuleName:                  "furya1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5",
		satypes.ModuleName:                           "furya1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6",
		clobmoduletypes.InsuranceFundName:            "furya1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq",
		rewardsmoduletypes.TreasuryAccountName:       "furya16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp",
		rewardsmoduletypes.VesterAccountName:         "furya1ltyc6y4skclzafvpznpt2qjwmfwgsndp458rmp",
		vestmoduletypes.CommunityTreasuryAccountName: "furya15ztc7xy42tn2ukkc0qjthkucw9ac63pgp70urn",
		vestmoduletypes.CommunityVesterAccountName:   "furya1wxje320an3karyc6mjw4zghs300dmrjkwn7xtk",
		delaymsgtypes.ModuleName:                     "furya1mkkvp26dngu6n8rmalaxyp3gwkjuzztq5zx6tr",
	}

	require.True(t, len(expectedModuleAccToAddresses) == len(app.GetMaccPerms()))
	for acc, address := range expectedModuleAccToAddresses {
		expectedAddr := authtypes.NewModuleAddress(acc).String()
		require.Equal(t, address, expectedAddr, "module (%v) should have address (%s)", acc, expectedAddr)
	}
}

func TestBlockedAddresses(t *testing.T) {
	expectedBlockedAddresses := map[string]bool{
		"furya17xpfvakm2amg962yls6f84z3kell8c5leqdyt2": true,
		"furya1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg": true,
		"furya1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605": true,
		"furya1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq": true,
		"furya1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5": true,
	}
	require.Equal(t, expectedBlockedAddresses, app.BlockedAddresses())
}

func TestMaccPerms(t *testing.T) {
	maccPerms := app.GetMaccPerms()
	expectedMaccPerms := map[string][]string{
		"bonded_tokens_pool":     {"burner", "staking"},
		"bridge":                 {"minter"},
		"distribution":           nil,
		"fee_collector":          nil,
		"gov":                    {"burner"},
		"insurance_fund":         nil,
		"not_bonded_tokens_pool": {"burner", "staking"},
		"subaccounts":            nil,
		"transfer":               {"minter", "burner"},
		"rewards_treasury":       nil,
		"rewards_vester":         nil,
		"community_treasury":     nil,
		"community_vester":       nil,
		"delaymsg":               nil,
	}
	require.Equal(t, expectedMaccPerms, maccPerms, "default macc perms list does not match expected")
}

func TestModuleAccountAddrs(t *testing.T) {
	expectedModuleAccAddresses := map[string]bool{
		"furya17xpfvakm2amg962yls6f84z3kell8c5leqdyt2": true, // x/auth.FeeCollector
		"furya1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv": true, // x/bridge
		"furya1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg": true, // x/distribution
		"furya1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq": true, // x/staking.bondedPool
		"furya1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605": true, // x/staking.notBondedPool
		"furya10d07y265gmmuvt4z0w9aw880jnsr700jnmapky": true, // x/ gov
		"furya1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5": true, // ibc transfer
		"furya1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6": true, // x/subaccount
		"furya1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq": true, // x/clob.insuranceFund
		"furya16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp": true, // x/rewards.treasury
		"furya1ltyc6y4skclzafvpznpt2qjwmfwgsndp458rmp": true, // x/rewards.vester
		"furya15ztc7xy42tn2ukkc0qjthkucw9ac63pgp70urn": true, // x/vest.communityTreasury
		"furya1wxje320an3karyc6mjw4zghs300dmrjkwn7xtk": true, // x/vest.communityVester
		"furya1mkkvp26dngu6n8rmalaxyp3gwkjuzztq5zx6tr": true, // x/delaymsg
	}

	require.Equal(t, expectedModuleAccAddresses, app.ModuleAccountAddrs())
}
