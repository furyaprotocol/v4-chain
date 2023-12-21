// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	lib "github.com/furyanprotocol/v4-chain/protocol/lib"
	delaymsgtypes "github.com/furyanprotocol/v4-chain/protocol/x/delaymsg/types"

	log "github.com/cometbft/cometbft/libs/log"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// DelayMsgKeeper is an autogenerated mock type for the DelayMsgKeeper type
type DelayMsgKeeper struct {
	mock.Mock
}

// DelayMessageByBlocks provides a mock function with given fields: ctx, msg, blockDelay
func (_m *DelayMsgKeeper) DelayMessageByBlocks(ctx types.Context, msg types.Msg, blockDelay uint32) (uint32, error) {
	ret := _m.Called(ctx, msg, blockDelay)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(types.Context, types.Msg, uint32) uint32); ok {
		r0 = rf(ctx, msg, blockDelay)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context, types.Msg, uint32) error); ok {
		r1 = rf(ctx, msg, blockDelay)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: ctx, id
func (_m *DelayMsgKeeper) DeleteMessage(ctx types.Context, id uint32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDelayedMessages provides a mock function with given fields: ctx
func (_m *DelayMsgKeeper) GetAllDelayedMessages(ctx types.Context) []*delaymsgtypes.DelayedMessage {
	ret := _m.Called(ctx)

	var r0 []*delaymsgtypes.DelayedMessage
	if rf, ok := ret.Get(0).(func(types.Context) []*delaymsgtypes.DelayedMessage); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*delaymsgtypes.DelayedMessage)
		}
	}

	return r0
}

// GetBlockMessageIds provides a mock function with given fields: ctx, blockHeight
func (_m *DelayMsgKeeper) GetBlockMessageIds(ctx types.Context, blockHeight uint32) (delaymsgtypes.BlockMessageIds, bool) {
	ret := _m.Called(ctx, blockHeight)

	var r0 delaymsgtypes.BlockMessageIds
	if rf, ok := ret.Get(0).(func(types.Context, uint32) delaymsgtypes.BlockMessageIds); ok {
		r0 = rf(ctx, blockHeight)
	} else {
		r0 = ret.Get(0).(delaymsgtypes.BlockMessageIds)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(types.Context, uint32) bool); ok {
		r1 = rf(ctx, blockHeight)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: ctx, id
func (_m *DelayMsgKeeper) GetMessage(ctx types.Context, id uint32) (delaymsgtypes.DelayedMessage, bool) {
	ret := _m.Called(ctx, id)

	var r0 delaymsgtypes.DelayedMessage
	if rf, ok := ret.Get(0).(func(types.Context, uint32) delaymsgtypes.DelayedMessage); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(delaymsgtypes.DelayedMessage)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(types.Context, uint32) bool); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// HasAuthority provides a mock function with given fields: authority
func (_m *DelayMsgKeeper) HasAuthority(authority string) bool {
	ret := _m.Called(authority)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(authority)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Logger provides a mock function with given fields: ctx
func (_m *DelayMsgKeeper) Logger(ctx types.Context) log.Logger {
	ret := _m.Called(ctx)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(types.Context) log.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// Router provides a mock function with given fields:
func (_m *DelayMsgKeeper) Router() lib.MsgRouter {
	ret := _m.Called()

	var r0 lib.MsgRouter
	if rf, ok := ret.Get(0).(func() lib.MsgRouter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(lib.MsgRouter)
		}
	}

	return r0
}

// SetDelayedMessage provides a mock function with given fields: ctx, msg
func (_m *DelayMsgKeeper) SetDelayedMessage(ctx types.Context, msg *delaymsgtypes.DelayedMessage) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *delaymsgtypes.DelayedMessage) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNextDelayedMessageId provides a mock function with given fields: ctx, nextDelayedMessageId
func (_m *DelayMsgKeeper) SetNextDelayedMessageId(ctx types.Context, nextDelayedMessageId uint32) {
	_m.Called(ctx, nextDelayedMessageId)
}

type mockConstructorTestingTNewDelayMsgKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewDelayMsgKeeper creates a new instance of DelayMsgKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDelayMsgKeeper(t mockConstructorTestingTNewDelayMsgKeeper) *DelayMsgKeeper {
	mock := &DelayMsgKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
