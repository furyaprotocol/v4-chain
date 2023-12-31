// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	bridgetypes "github.com/furyanprotocol/v4-chain/protocol/x/bridge/types"
	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PrepareBridgeKeeper is an autogenerated mock type for the PrepareBridgeKeeper type
type PrepareBridgeKeeper struct {
	mock.Mock
}

// GetAcknowledgeBridges provides a mock function with given fields: ctx, blockTimestamp
func (_m *PrepareBridgeKeeper) GetAcknowledgeBridges(ctx types.Context, blockTimestamp time.Time) *bridgetypes.MsgAcknowledgeBridges {
	ret := _m.Called(ctx, blockTimestamp)

	var r0 *bridgetypes.MsgAcknowledgeBridges
	if rf, ok := ret.Get(0).(func(types.Context, time.Time) *bridgetypes.MsgAcknowledgeBridges); ok {
		r0 = rf(ctx, blockTimestamp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridgetypes.MsgAcknowledgeBridges)
		}
	}

	return r0
}

type mockConstructorTestingTNewPrepareBridgeKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewPrepareBridgeKeeper creates a new instance of PrepareBridgeKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPrepareBridgeKeeper(t mockConstructorTestingTNewPrepareBridgeKeeper) *PrepareBridgeKeeper {
	mock := &PrepareBridgeKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
