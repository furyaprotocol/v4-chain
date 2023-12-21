package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/furyanprotocol/v4-chain/protocol/lib"
	blocktime "github.com/furyanprotocol/v4-chain/protocol/x/blocktime/types"
	bridge "github.com/furyanprotocol/v4-chain/protocol/x/bridge/types"
	clob "github.com/furyanprotocol/v4-chain/protocol/x/clob/types"
	delaymsg "github.com/furyanprotocol/v4-chain/protocol/x/delaymsg/types"
	feetiers "github.com/furyanprotocol/v4-chain/protocol/x/feetiers/types"
	perpetuals "github.com/furyanprotocol/v4-chain/protocol/x/perpetuals/types"
	prices "github.com/furyanprotocol/v4-chain/protocol/x/prices/types"
	rewards "github.com/furyanprotocol/v4-chain/protocol/x/rewards/types"
	sending "github.com/furyanprotocol/v4-chain/protocol/x/sending/types"
	stats "github.com/furyanprotocol/v4-chain/protocol/x/stats/types"
	vest "github.com/furyanprotocol/v4-chain/protocol/x/vest/types"
)

var (
	// InternalMsgSamplesAll are msgs that are used only used internally.
	InternalMsgSamplesAll = lib.MergeAllMapsMustHaveDistinctKeys(InternalMsgSamplesGovAuth)

	// InternalMsgSamplesGovAuth are msgs that are used only used internally.
	// GovAuth means that these messages must originate from the gov module and
	// signed by gov module account.
	InternalMsgSamplesGovAuth = map[string]sdk.Msg{
		// ------- CosmosSDK default modules
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams": &auth.MsgUpdateParams{},

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         &bank.MsgSetSendEnabled{},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": nil,
		"/cosmos.bank.v1beta1.MsgUpdateParams":           &bank.MsgUpdateParams{},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   nil,

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         &consensus.MsgUpdateParams{},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": nil,

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":         &crisis.MsgUpdateParams{},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse": nil,

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":         &distribution.MsgCommunityPoolSpend{},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse": nil,
		"/cosmos.distribution.v1beta1.MsgUpdateParams":               &distribution.MsgUpdateParams{},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":       nil,

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent":         &gov.MsgExecLegacyContent{},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse": nil,
		"/cosmos.gov.v1.MsgUpdateParams":              &gov.MsgUpdateParams{},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":      nil,

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         &slashing.MsgUpdateParams{},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": nil,

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams":         &staking.MsgUpdateParams{},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse": nil,

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":           &upgrade.MsgCancelUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":   nil,
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":         &upgrade.MsgSoftwareUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse": nil,

		// ------- Custom modules
		// blocktime
		"/furyaprotocol.blocktime.MsgUpdateDowntimeParams":         &blocktime.MsgUpdateDowntimeParams{},
		"/furyaprotocol.blocktime.MsgUpdateDowntimeParamsResponse": nil,

		// bridge
		"/furyaprotocol.bridge.MsgCompleteBridge":              &bridge.MsgCompleteBridge{},
		"/furyaprotocol.bridge.MsgCompleteBridgeResponse":      nil,
		"/furyaprotocol.bridge.MsgUpdateEventParams":           &bridge.MsgUpdateEventParams{},
		"/furyaprotocol.bridge.MsgUpdateEventParamsResponse":   nil,
		"/furyaprotocol.bridge.MsgUpdateProposeParams":         &bridge.MsgUpdateProposeParams{},
		"/furyaprotocol.bridge.MsgUpdateProposeParamsResponse": nil,
		"/furyaprotocol.bridge.MsgUpdateSafetyParams":          &bridge.MsgUpdateSafetyParams{},
		"/furyaprotocol.bridge.MsgUpdateSafetyParamsResponse":  nil,

		// clob
		"/furyaprotocol.clob.MsgCreateClobPair":                             &clob.MsgCreateClobPair{},
		"/furyaprotocol.clob.MsgCreateClobPairResponse":                     nil,
		"/furyaprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          &clob.MsgUpdateBlockRateLimitConfiguration{},
		"/furyaprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  nil,
		"/furyaprotocol.clob.MsgUpdateClobPair":                             &clob.MsgUpdateClobPair{},
		"/furyaprotocol.clob.MsgUpdateClobPairResponse":                     nil,
		"/furyaprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         &clob.MsgUpdateEquityTierLimitConfiguration{},
		"/furyaprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": nil,
		"/furyaprotocol.clob.MsgUpdateLiquidationsConfig":                   &clob.MsgUpdateLiquidationsConfig{},
		"/furyaprotocol.clob.MsgUpdateLiquidationsConfigResponse":           nil,

		// delaymsg
		"/furyaprotocol.delaymsg.MsgDelayMessage":         &delaymsg.MsgDelayMessage{},
		"/furyaprotocol.delaymsg.MsgDelayMessageResponse": nil,

		// feetiers
		"/furyaprotocol.feetiers.MsgUpdatePerpetualFeeParams":         &feetiers.MsgUpdatePerpetualFeeParams{},
		"/furyaprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": nil,

		// perpetuals
		"/furyaprotocol.perpetuals.MsgCreatePerpetual":               &perpetuals.MsgCreatePerpetual{},
		"/furyaprotocol.perpetuals.MsgCreatePerpetualResponse":       nil,
		"/furyaprotocol.perpetuals.MsgSetLiquidityTier":              &perpetuals.MsgSetLiquidityTier{},
		"/furyaprotocol.perpetuals.MsgSetLiquidityTierResponse":      nil,
		"/furyaprotocol.perpetuals.MsgUpdateParams":                  &perpetuals.MsgUpdateParams{},
		"/furyaprotocol.perpetuals.MsgUpdateParamsResponse":          nil,
		"/furyaprotocol.perpetuals.MsgUpdatePerpetualParams":         &perpetuals.MsgUpdatePerpetualParams{},
		"/furyaprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": nil,

		// prices
		"/furyaprotocol.prices.MsgCreateOracleMarket":         &prices.MsgCreateOracleMarket{},
		"/furyaprotocol.prices.MsgCreateOracleMarketResponse": nil,
		"/furyaprotocol.prices.MsgUpdateMarketParam":          &prices.MsgUpdateMarketParam{},
		"/furyaprotocol.prices.MsgUpdateMarketParamResponse":  nil,

		// rewards
		"/furyaprotocol.rewards.MsgUpdateParams":         &rewards.MsgUpdateParams{},
		"/furyaprotocol.rewards.MsgUpdateParamsResponse": nil,

		// sending
		"/furyaprotocol.sending.MsgSendFromModuleToAccount":         &sending.MsgSendFromModuleToAccount{},
		"/furyaprotocol.sending.MsgSendFromModuleToAccountResponse": nil,

		// stats
		"/furyaprotocol.stats.MsgUpdateParams":         &stats.MsgUpdateParams{},
		"/furyaprotocol.stats.MsgUpdateParamsResponse": nil,

		// vest
		"/furyaprotocol.vest.MsgSetVestEntry":            &vest.MsgSetVestEntry{},
		"/furyaprotocol.vest.MsgSetVestEntryResponse":    nil,
		"/furyaprotocol.vest.MsgDeleteVestEntry":         &vest.MsgDeleteVestEntry{},
		"/furyaprotocol.vest.MsgDeleteVestEntryResponse": nil,
	}
)
