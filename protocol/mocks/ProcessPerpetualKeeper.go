// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	perpetualstypes "github.com/furyanprotocol/v4-chain/protocol/x/perpetuals/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// ProcessPerpetualKeeper is an autogenerated mock type for the ProcessPerpetualKeeper type
type ProcessPerpetualKeeper struct {
	mock.Mock
}

// GetPerpetual provides a mock function with given fields: ctx, id
func (_m *ProcessPerpetualKeeper) GetPerpetual(ctx types.Context, id uint32) (perpetualstypes.Perpetual, error) {
	ret := _m.Called(ctx, id)

	var r0 perpetualstypes.Perpetual
	if rf, ok := ret.Get(0).(func(types.Context, uint32) perpetualstypes.Perpetual); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(perpetualstypes.Perpetual)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSettlementPpm provides a mock function with given fields: ctx, perpetualId, quantums, index
func (_m *ProcessPerpetualKeeper) GetSettlementPpm(ctx types.Context, perpetualId uint32, quantums *big.Int, index *big.Int) (*big.Int, *big.Int, error) {
	ret := _m.Called(ctx, perpetualId, quantums, index)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(ctx, perpetualId, quantums, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(types.Context, uint32, *big.Int, *big.Int) *big.Int); ok {
		r1 = rf(ctx, perpetualId, quantums, index)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(types.Context, uint32, *big.Int, *big.Int) error); ok {
		r2 = rf(ctx, perpetualId, quantums, index)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MaybeProcessNewFundingTickEpoch provides a mock function with given fields: ctx
func (_m *ProcessPerpetualKeeper) MaybeProcessNewFundingTickEpoch(ctx types.Context) {
	_m.Called(ctx)
}

type mockConstructorTestingTNewProcessPerpetualKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewProcessPerpetualKeeper creates a new instance of ProcessPerpetualKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProcessPerpetualKeeper(t mockConstructorTestingTNewProcessPerpetualKeeper) *ProcessPerpetualKeeper {
	mock := &ProcessPerpetualKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
