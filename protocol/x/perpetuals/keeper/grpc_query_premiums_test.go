package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/furyanprotocol/v4-chain/protocol/testutil/keeper"
	"github.com/furyanprotocol/v4-chain/protocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestPremiumVotes(t *testing.T) {
	pc := keepertest.PerpetualsKeepers(t)
	wctx := sdk.WrapSDKContext(pc.Ctx)

	tests := map[string]struct {
		req         *types.QueryPremiumVotesRequest
		res         *types.QueryPremiumVotesResponse
		expectedErr error
	}{
		"nil request": {
			req:         nil,
			res:         nil,
			expectedErr: status.Error(codes.InvalidArgument, "invalid request"),
		},
		"valid request": {
			req: &types.QueryPremiumVotesRequest{},
			res: &types.QueryPremiumVotesResponse{
				PremiumVotes: types.PremiumStore{},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := pc.PerpetualsKeeper.PremiumVotes(wctx, tc.req)
			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.Equal(t, tc.res, res)
			}
		})
	}
}

func TestPremiumSamples(t *testing.T) {
	pc := keepertest.PerpetualsKeepers(t)
	wctx := sdk.WrapSDKContext(pc.Ctx)

	tests := map[string]struct {
		req         *types.QueryPremiumSamplesRequest
		res         *types.QueryPremiumSamplesResponse
		expectedErr error
	}{
		"nil request": {
			req:         nil,
			res:         nil,
			expectedErr: status.Error(codes.InvalidArgument, "invalid request"),
		},
		"valid request": {
			req: &types.QueryPremiumSamplesRequest{},
			res: &types.QueryPremiumSamplesResponse{
				PremiumSamples: types.PremiumStore{},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := pc.PerpetualsKeeper.PremiumSamples(wctx, tc.req)
			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.Equal(t, tc.res, res)
			}
		})
	}
}
