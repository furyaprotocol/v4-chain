package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyanprotocol/v4-chain/protocol/x/bridge/types"
)

// UpdateEventParams updates the EventParams in state.
func (k msgServer) UpdateEventParams(
	goCtx context.Context,
	msg *types.MsgUpdateEventParams,
) (*types.MsgUpdateEventParamsResponse, error) {
	if !k.Keeper.HasAuthority(msg.GetAuthority()) {
		return nil, errors.Wrapf(
			types.ErrInvalidAuthority,
			"message authority %s is not valid for sending update event params messages",
			msg.GetAuthority(),
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.UpdateEventParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateEventParamsResponse{}, nil
}
