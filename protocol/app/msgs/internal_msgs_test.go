package msgs_test

import (
	"sort"
	"testing"

	"github.com/furyanprotocol/v4-chain/protocol/app/msgs"
	"github.com/furyanprotocol/v4-chain/protocol/lib"
	"github.com/stretchr/testify/require"
)

func TestInternalMsgSamples_All_Key(t *testing.T) {
	expectedAllInternalMsgs := lib.MergeAllMapsMustHaveDistinctKeys(msgs.InternalMsgSamplesGovAuth)
	require.Equal(t, expectedAllInternalMsgs, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_All_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_Gov_Key(t *testing.T) {
	expectedMsgs := []string{
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams",

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled",
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse",
		"/cosmos.bank.v1beta1.MsgUpdateParams",
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse",

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams",
		"/cosmos.consensus.v1.MsgUpdateParamsResponse",

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams",
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse",

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse",
		"/cosmos.distribution.v1beta1.MsgUpdateParams",
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse",

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent",
		"/cosmos.gov.v1.MsgExecLegacyContentResponse",
		"/cosmos.gov.v1.MsgUpdateParams",
		"/cosmos.gov.v1.MsgUpdateParamsResponse",

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams",
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse",

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams",
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse",

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse",

		// blocktime
		"/furyaprotocol.blocktime.MsgUpdateDowntimeParams",
		"/furyaprotocol.blocktime.MsgUpdateDowntimeParamsResponse",

		// bridge
		"/furyaprotocol.bridge.MsgCompleteBridge",
		"/furyaprotocol.bridge.MsgCompleteBridgeResponse",
		"/furyaprotocol.bridge.MsgUpdateEventParams",
		"/furyaprotocol.bridge.MsgUpdateEventParamsResponse",
		"/furyaprotocol.bridge.MsgUpdateProposeParams",
		"/furyaprotocol.bridge.MsgUpdateProposeParamsResponse",
		"/furyaprotocol.bridge.MsgUpdateSafetyParams",
		"/furyaprotocol.bridge.MsgUpdateSafetyParamsResponse",

		// clob
		"/furyaprotocol.clob.MsgCreateClobPair",
		"/furyaprotocol.clob.MsgCreateClobPairResponse",
		"/furyaprotocol.clob.MsgUpdateBlockRateLimitConfiguration",
		"/furyaprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse",
		"/furyaprotocol.clob.MsgUpdateClobPair",
		"/furyaprotocol.clob.MsgUpdateClobPairResponse",
		"/furyaprotocol.clob.MsgUpdateEquityTierLimitConfiguration",
		"/furyaprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse",
		"/furyaprotocol.clob.MsgUpdateLiquidationsConfig",
		"/furyaprotocol.clob.MsgUpdateLiquidationsConfigResponse",

		// delaymsg
		"/furyaprotocol.delaymsg.MsgDelayMessage",
		"/furyaprotocol.delaymsg.MsgDelayMessageResponse",

		// feetiers
		"/furyaprotocol.feetiers.MsgUpdatePerpetualFeeParams",
		"/furyaprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse",

		// perpeutals
		"/furyaprotocol.perpetuals.MsgCreatePerpetual",
		"/furyaprotocol.perpetuals.MsgCreatePerpetualResponse",
		"/furyaprotocol.perpetuals.MsgSetLiquidityTier",
		"/furyaprotocol.perpetuals.MsgSetLiquidityTierResponse",
		"/furyaprotocol.perpetuals.MsgUpdateParams",
		"/furyaprotocol.perpetuals.MsgUpdateParamsResponse",
		"/furyaprotocol.perpetuals.MsgUpdatePerpetualParams",
		"/furyaprotocol.perpetuals.MsgUpdatePerpetualParamsResponse",

		// prices
		"/furyaprotocol.prices.MsgCreateOracleMarket",
		"/furyaprotocol.prices.MsgCreateOracleMarketResponse",
		"/furyaprotocol.prices.MsgUpdateMarketParam",
		"/furyaprotocol.prices.MsgUpdateMarketParamResponse",

		// rewards
		"/furyaprotocol.rewards.MsgUpdateParams",
		"/furyaprotocol.rewards.MsgUpdateParamsResponse",

		// sending
		"/furyaprotocol.sending.MsgSendFromModuleToAccount",
		"/furyaprotocol.sending.MsgSendFromModuleToAccountResponse",

		// stats
		"/furyaprotocol.stats.MsgUpdateParams",
		"/furyaprotocol.stats.MsgUpdateParamsResponse",

		// vest
		"/furyaprotocol.vest.MsgDeleteVestEntry",
		"/furyaprotocol.vest.MsgDeleteVestEntryResponse",
		"/furyaprotocol.vest.MsgSetVestEntry",
		"/furyaprotocol.vest.MsgSetVestEntryResponse",
	}

	require.Equal(t, expectedMsgs, lib.GetSortedKeys[sort.StringSlice](msgs.InternalMsgSamplesGovAuth))
}

func TestInternalMsgSamples_Gov_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesGovAuth)
}
