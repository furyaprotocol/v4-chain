package msgs

import (
	"github.com/furyanprotocol/v4-chain/protocol/lib"
)

var (
	// AllTypeMessages is a list of all messages and types that are used in the app.
	// This list comes from the app's `InterfaceRegistry`.
	AllTypeMessages = map[string]struct{}{
		// auth
		"/cosmos.auth.v1beta1.BaseAccount":      {},
		"/cosmos.auth.v1beta1.ModuleAccount":    {},
		"/cosmos.auth.v1beta1.ModuleCredential": {},
		"/cosmos.auth.v1beta1.MsgUpdateParams":  {},

		// bank
		"/cosmos.bank.v1beta1.MsgMultiSend":              {},
		"/cosmos.bank.v1beta1.MsgMultiSendResponse":      {},
		"/cosmos.bank.v1beta1.MsgSend":                   {},
		"/cosmos.bank.v1beta1.MsgSendResponse":           {},
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         {},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": {},
		"/cosmos.bank.v1beta1.MsgUpdateParams":           {},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   {},
		"/cosmos.bank.v1beta1.SendAuthorization":         {},
		"/cosmos.bank.v1beta1.Supply":                    {},

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         {},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": {},

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":            {},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse":    {},
		"/cosmos.crisis.v1beta1.MsgVerifyInvariant":         {},
		"/cosmos.crisis.v1beta1.MsgVerifyInvariantResponse": {},

		// crypto
		"/cosmos.crypto.ed25519.PrivKey":            {},
		"/cosmos.crypto.ed25519.PubKey":             {},
		"/cosmos.crypto.multisig.LegacyAminoPubKey": {},
		"/cosmos.crypto.secp256k1.PrivKey":          {},
		"/cosmos.crypto.secp256k1.PubKey":           {},
		"/cosmos.crypto.secp256r1.PubKey":           {},

		// distribution
		"/cosmos.distribution.v1beta1.CommunityPoolSpendProposal":             {},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":                  {},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse":          {},
		"/cosmos.distribution.v1beta1.MsgFundCommunityPool":                   {},
		"/cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse":           {},
		"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":                  {},
		"/cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse":          {},
		"/cosmos.distribution.v1beta1.MsgUpdateParams":                        {},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":                {},
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":             {},
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse":     {},
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":         {},
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse": {},

		// evidence
		"/cosmos.evidence.v1beta1.Equivocation":              {},
		"/cosmos.evidence.v1beta1.MsgSubmitEvidence":         {},
		"/cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse": {},

		// feegrant
		"/cosmos.feegrant.v1beta1.AllowedMsgAllowance":        {},
		"/cosmos.feegrant.v1beta1.BasicAllowance":             {},
		"/cosmos.feegrant.v1beta1.MsgGrantAllowance":          {},
		"/cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse":  {},
		"/cosmos.feegrant.v1beta1.MsgRevokeAllowance":         {},
		"/cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse": {},
		"/cosmos.feegrant.v1beta1.PeriodicAllowance":          {},

		// gov
		"/cosmos.gov.v1.MsgDeposit":                     {},
		"/cosmos.gov.v1.MsgDepositResponse":             {},
		"/cosmos.gov.v1.MsgExecLegacyContent":           {},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse":   {},
		"/cosmos.gov.v1.MsgSubmitProposal":              {},
		"/cosmos.gov.v1.MsgSubmitProposalResponse":      {},
		"/cosmos.gov.v1.MsgUpdateParams":                {},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":        {},
		"/cosmos.gov.v1.MsgVote":                        {},
		"/cosmos.gov.v1.MsgVoteResponse":                {},
		"/cosmos.gov.v1.MsgVoteWeighted":                {},
		"/cosmos.gov.v1.MsgVoteWeightedResponse":        {},
		"/cosmos.gov.v1beta1.MsgDeposit":                {},
		"/cosmos.gov.v1beta1.MsgDepositResponse":        {},
		"/cosmos.gov.v1beta1.MsgSubmitProposal":         {},
		"/cosmos.gov.v1beta1.MsgSubmitProposalResponse": {},
		"/cosmos.gov.v1beta1.MsgVote":                   {},
		"/cosmos.gov.v1beta1.MsgVoteResponse":           {},
		"/cosmos.gov.v1beta1.MsgVoteWeighted":           {},
		"/cosmos.gov.v1beta1.MsgVoteWeightedResponse":   {},
		"/cosmos.gov.v1beta1.TextProposal":              {},

		// params
		"/cosmos.params.v1beta1.ParameterChangeProposal": {},

		// slashing
		"/cosmos.slashing.v1beta1.MsgUnjail":               {},
		"/cosmos.slashing.v1beta1.MsgUnjailResponse":       {},
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         {},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": {},

		// staking
		"/cosmos.staking.v1beta1.MsgBeginRedelegate":                   {},
		"/cosmos.staking.v1beta1.MsgBeginRedelegateResponse":           {},
		"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation":         {},
		"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse": {},
		"/cosmos.staking.v1beta1.MsgCreateValidator":                   {},
		"/cosmos.staking.v1beta1.MsgCreateValidatorResponse":           {},
		"/cosmos.staking.v1beta1.MsgDelegate":                          {},
		"/cosmos.staking.v1beta1.MsgDelegateResponse":                  {},
		"/cosmos.staking.v1beta1.MsgEditValidator":                     {},
		"/cosmos.staking.v1beta1.MsgEditValidatorResponse":             {},
		"/cosmos.staking.v1beta1.MsgUndelegate":                        {},
		"/cosmos.staking.v1beta1.MsgUndelegateResponse":                {},
		"/cosmos.staking.v1beta1.MsgUpdateParams":                      {},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse":              {},
		"/cosmos.staking.v1beta1.StakeAuthorization":                   {},

		// tx
		"/cosmos.tx.v1beta1.Tx": {},

		// upgrade
		"/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal": {},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":              {},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":      {},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":            {},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse":    {},
		"/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal":       {},

		// blocktime
		"/furyaprotocol.blocktime.MsgUpdateDowntimeParams":         {},
		"/furyaprotocol.blocktime.MsgUpdateDowntimeParamsResponse": {},

		// bridge
		"/furyaprotocol.bridge.MsgAcknowledgeBridges":          {},
		"/furyaprotocol.bridge.MsgAcknowledgeBridgesResponse":  {},
		"/furyaprotocol.bridge.MsgCompleteBridge":              {},
		"/furyaprotocol.bridge.MsgCompleteBridgeResponse":      {},
		"/furyaprotocol.bridge.MsgUpdateEventParams":           {},
		"/furyaprotocol.bridge.MsgUpdateEventParamsResponse":   {},
		"/furyaprotocol.bridge.MsgUpdateProposeParams":         {},
		"/furyaprotocol.bridge.MsgUpdateProposeParamsResponse": {},
		"/furyaprotocol.bridge.MsgUpdateSafetyParams":          {},
		"/furyaprotocol.bridge.MsgUpdateSafetyParamsResponse":  {},

		// clob
		"/furyaprotocol.clob.MsgCancelOrder":                                {},
		"/furyaprotocol.clob.MsgCancelOrderResponse":                        {},
		"/furyaprotocol.clob.MsgCreateClobPair":                             {},
		"/furyaprotocol.clob.MsgCreateClobPairResponse":                     {},
		"/furyaprotocol.clob.MsgPlaceOrder":                                 {},
		"/furyaprotocol.clob.MsgPlaceOrderResponse":                         {},
		"/furyaprotocol.clob.MsgProposedOperations":                         {},
		"/furyaprotocol.clob.MsgProposedOperationsResponse":                 {},
		"/furyaprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          {},
		"/furyaprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  {},
		"/furyaprotocol.clob.MsgUpdateClobPair":                             {},
		"/furyaprotocol.clob.MsgUpdateClobPairResponse":                     {},
		"/furyaprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         {},
		"/furyaprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": {},
		"/furyaprotocol.clob.MsgUpdateLiquidationsConfig":                   {},
		"/furyaprotocol.clob.MsgUpdateLiquidationsConfigResponse":           {},

		// delaymsg
		"/furyaprotocol.delaymsg.MsgDelayMessage":         {},
		"/furyaprotocol.delaymsg.MsgDelayMessageResponse": {},

		// feetiers
		"/furyaprotocol.feetiers.MsgUpdatePerpetualFeeParams":         {},
		"/furyaprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": {},

		// perpetuals
		"/furyaprotocol.perpetuals.MsgAddPremiumVotes":               {},
		"/furyaprotocol.perpetuals.MsgAddPremiumVotesResponse":       {},
		"/furyaprotocol.perpetuals.MsgCreatePerpetual":               {},
		"/furyaprotocol.perpetuals.MsgCreatePerpetualResponse":       {},
		"/furyaprotocol.perpetuals.MsgSetLiquidityTier":              {},
		"/furyaprotocol.perpetuals.MsgSetLiquidityTierResponse":      {},
		"/furyaprotocol.perpetuals.MsgUpdateParams":                  {},
		"/furyaprotocol.perpetuals.MsgUpdateParamsResponse":          {},
		"/furyaprotocol.perpetuals.MsgUpdatePerpetualParams":         {},
		"/furyaprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": {},

		// prices
		"/furyaprotocol.prices.MsgCreateOracleMarket":         {},
		"/furyaprotocol.prices.MsgCreateOracleMarketResponse": {},
		"/furyaprotocol.prices.MsgUpdateMarketPrices":         {},
		"/furyaprotocol.prices.MsgUpdateMarketPricesResponse": {},
		"/furyaprotocol.prices.MsgUpdateMarketParam":          {},
		"/furyaprotocol.prices.MsgUpdateMarketParamResponse":  {},

		// sending
		"/furyaprotocol.sending.MsgCreateTransfer":                  {},
		"/furyaprotocol.sending.MsgCreateTransferResponse":          {},
		"/furyaprotocol.sending.MsgDepositToSubaccount":             {},
		"/furyaprotocol.sending.MsgDepositToSubaccountResponse":     {},
		"/furyaprotocol.sending.MsgWithdrawFromSubaccount":          {},
		"/furyaprotocol.sending.MsgWithdrawFromSubaccountResponse":  {},
		"/furyaprotocol.sending.MsgSendFromModuleToAccount":         {},
		"/furyaprotocol.sending.MsgSendFromModuleToAccountResponse": {},

		// stats
		"/furyaprotocol.stats.MsgUpdateParams":         {},
		"/furyaprotocol.stats.MsgUpdateParamsResponse": {},

		// vest
		"/furyaprotocol.vest.MsgSetVestEntry":            {},
		"/furyaprotocol.vest.MsgSetVestEntryResponse":    {},
		"/furyaprotocol.vest.MsgDeleteVestEntry":         {},
		"/furyaprotocol.vest.MsgDeleteVestEntryResponse": {},

		// rewards
		"/furyaprotocol.rewards.MsgUpdateParams":         {},
		"/furyaprotocol.rewards.MsgUpdateParamsResponse": {},

		// ibc.applications
		"/ibc.applications.transfer.v1.MsgTransfer":           {},
		"/ibc.applications.transfer.v1.MsgTransferResponse":   {},
		"/ibc.applications.transfer.v1.TransferAuthorization": {},

		// ibc.core.channel
		"/ibc.core.channel.v1.Channel":                        {},
		"/ibc.core.channel.v1.Counterparty":                   {},
		"/ibc.core.channel.v1.MsgAcknowledgement":             {},
		"/ibc.core.channel.v1.MsgAcknowledgementResponse":     {},
		"/ibc.core.channel.v1.MsgChannelCloseConfirm":         {},
		"/ibc.core.channel.v1.MsgChannelCloseConfirmResponse": {},
		"/ibc.core.channel.v1.MsgChannelCloseInit":            {},
		"/ibc.core.channel.v1.MsgChannelCloseInitResponse":    {},
		"/ibc.core.channel.v1.MsgChannelOpenAck":              {},
		"/ibc.core.channel.v1.MsgChannelOpenAckResponse":      {},
		"/ibc.core.channel.v1.MsgChannelOpenConfirm":          {},
		"/ibc.core.channel.v1.MsgChannelOpenConfirmResponse":  {},
		"/ibc.core.channel.v1.MsgChannelOpenInit":             {},
		"/ibc.core.channel.v1.MsgChannelOpenInitResponse":     {},
		"/ibc.core.channel.v1.MsgChannelOpenTry":              {},
		"/ibc.core.channel.v1.MsgChannelOpenTryResponse":      {},
		"/ibc.core.channel.v1.MsgRecvPacket":                  {},
		"/ibc.core.channel.v1.MsgRecvPacketResponse":          {},
		"/ibc.core.channel.v1.MsgTimeout":                     {},
		"/ibc.core.channel.v1.MsgTimeoutOnClose":              {},
		"/ibc.core.channel.v1.MsgTimeoutOnCloseResponse":      {},
		"/ibc.core.channel.v1.MsgTimeoutResponse":             {},
		"/ibc.core.channel.v1.Packet":                         {},

		// ibc.core.client
		"/ibc.core.client.v1.ClientUpdateProposal":          {},
		"/ibc.core.client.v1.Height":                        {},
		"/ibc.core.client.v1.MsgCreateClient":               {},
		"/ibc.core.client.v1.MsgCreateClientResponse":       {},
		"/ibc.core.client.v1.MsgSubmitMisbehaviour":         {},
		"/ibc.core.client.v1.MsgSubmitMisbehaviourResponse": {},
		"/ibc.core.client.v1.MsgUpdateClient":               {},
		"/ibc.core.client.v1.MsgUpdateClientResponse":       {},
		"/ibc.core.client.v1.MsgUpgradeClient":              {},
		"/ibc.core.client.v1.MsgUpgradeClientResponse":      {},
		"/ibc.core.client.v1.UpgradeProposal":               {},

		// ibc.core.commitment
		"/ibc.core.commitment.v1.MerklePath":   {},
		"/ibc.core.commitment.v1.MerklePrefix": {},
		"/ibc.core.commitment.v1.MerkleProof":  {},
		"/ibc.core.commitment.v1.MerkleRoot":   {},

		// ibc.core.connection
		"/ibc.core.connection.v1.ConnectionEnd":                    {},
		"/ibc.core.connection.v1.Counterparty":                     {},
		"/ibc.core.connection.v1.MsgConnectionOpenAck":             {},
		"/ibc.core.connection.v1.MsgConnectionOpenAckResponse":     {},
		"/ibc.core.connection.v1.MsgConnectionOpenConfirm":         {},
		"/ibc.core.connection.v1.MsgConnectionOpenConfirmResponse": {},
		"/ibc.core.connection.v1.MsgConnectionOpenInit":            {},
		"/ibc.core.connection.v1.MsgConnectionOpenInitResponse":    {},
		"/ibc.core.connection.v1.MsgConnectionOpenTry":             {},
		"/ibc.core.connection.v1.MsgConnectionOpenTryResponse":     {},
		"/ibc.core.connection.v1.Version":                          {},

		// ibc.lightclients
		"/ibc.lightclients.localhost.v2.ClientState":     {},
		"/ibc.lightclients.tendermint.v1.ClientState":    {},
		"/ibc.lightclients.tendermint.v1.ConsensusState": {},
		"/ibc.lightclients.tendermint.v1.Header":         {},
		"/ibc.lightclients.tendermint.v1.Misbehaviour":   {},
	}

	// DisallowMsgs are messages that cannot be externally submitted.
	DisallowMsgs = lib.MergeAllMapsMustHaveDistinctKeys(
		AppInjectedMsgSamples,
		InternalMsgSamplesAll,
		NestedMsgSamples,
		UnsupportedMsgSamples,
	)

	// AllowMsgs are messages that can be externally submitted.
	AllowMsgs = NormalMsgs
)
