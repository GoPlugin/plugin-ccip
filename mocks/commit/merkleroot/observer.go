// Code generated by mockery v2.43.2. DO NOT EDIT.

package merkleroot

import (
	context "context"

	ccipocr3 "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"

	mock "github.com/stretchr/testify/mock"

	plugintypes "github.com/goplugin/plugin-ccip/internal/plugintypes"

	types "github.com/goplugin/plugin-ccip/commit/merkleroot/rmn/types"
)

// MockObserver is an autogenerated mock type for the Observer type
type MockObserver struct {
	mock.Mock
}

type MockObserver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObserver) EXPECT() *MockObserver_Expecter {
	return &MockObserver_Expecter{mock: &_m.Mock}
}

// ObserveFChain provides a mock function with given fields:
func (_m *MockObserver) ObserveFChain() map[ccipocr3.ChainSelector]int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ObserveFChain")
	}

	var r0 map[ccipocr3.ChainSelector]int
	if rf, ok := ret.Get(0).(func() map[ccipocr3.ChainSelector]int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ccipocr3.ChainSelector]int)
		}
	}

	return r0
}

// MockObserver_ObserveFChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObserveFChain'
type MockObserver_ObserveFChain_Call struct {
	*mock.Call
}

// ObserveFChain is a helper method to define mock.On call
func (_e *MockObserver_Expecter) ObserveFChain() *MockObserver_ObserveFChain_Call {
	return &MockObserver_ObserveFChain_Call{Call: _e.mock.On("ObserveFChain")}
}

func (_c *MockObserver_ObserveFChain_Call) Run(run func()) *MockObserver_ObserveFChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockObserver_ObserveFChain_Call) Return(_a0 map[ccipocr3.ChainSelector]int) *MockObserver_ObserveFChain_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObserver_ObserveFChain_Call) RunAndReturn(run func() map[ccipocr3.ChainSelector]int) *MockObserver_ObserveFChain_Call {
	_c.Call.Return(run)
	return _c
}

// ObserveLatestOnRampSeqNums provides a mock function with given fields: ctx, destChain
func (_m *MockObserver) ObserveLatestOnRampSeqNums(ctx context.Context, destChain ccipocr3.ChainSelector) []plugintypes.SeqNumChain {
	ret := _m.Called(ctx, destChain)

	if len(ret) == 0 {
		panic("no return value specified for ObserveLatestOnRampSeqNums")
	}

	var r0 []plugintypes.SeqNumChain
	if rf, ok := ret.Get(0).(func(context.Context, ccipocr3.ChainSelector) []plugintypes.SeqNumChain); ok {
		r0 = rf(ctx, destChain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]plugintypes.SeqNumChain)
		}
	}

	return r0
}

// MockObserver_ObserveLatestOnRampSeqNums_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObserveLatestOnRampSeqNums'
type MockObserver_ObserveLatestOnRampSeqNums_Call struct {
	*mock.Call
}

// ObserveLatestOnRampSeqNums is a helper method to define mock.On call
//   - ctx context.Context
//   - destChain ccipocr3.ChainSelector
func (_e *MockObserver_Expecter) ObserveLatestOnRampSeqNums(ctx interface{}, destChain interface{}) *MockObserver_ObserveLatestOnRampSeqNums_Call {
	return &MockObserver_ObserveLatestOnRampSeqNums_Call{Call: _e.mock.On("ObserveLatestOnRampSeqNums", ctx, destChain)}
}

func (_c *MockObserver_ObserveLatestOnRampSeqNums_Call) Run(run func(ctx context.Context, destChain ccipocr3.ChainSelector)) *MockObserver_ObserveLatestOnRampSeqNums_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockObserver_ObserveLatestOnRampSeqNums_Call) Return(_a0 []plugintypes.SeqNumChain) *MockObserver_ObserveLatestOnRampSeqNums_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObserver_ObserveLatestOnRampSeqNums_Call) RunAndReturn(run func(context.Context, ccipocr3.ChainSelector) []plugintypes.SeqNumChain) *MockObserver_ObserveLatestOnRampSeqNums_Call {
	_c.Call.Return(run)
	return _c
}

// ObserveMerkleRoots provides a mock function with given fields: ctx, ranges
func (_m *MockObserver) ObserveMerkleRoots(ctx context.Context, ranges []plugintypes.ChainRange) []ccipocr3.MerkleRootChain {
	ret := _m.Called(ctx, ranges)

	if len(ret) == 0 {
		panic("no return value specified for ObserveMerkleRoots")
	}

	var r0 []ccipocr3.MerkleRootChain
	if rf, ok := ret.Get(0).(func(context.Context, []plugintypes.ChainRange) []ccipocr3.MerkleRootChain); ok {
		r0 = rf(ctx, ranges)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipocr3.MerkleRootChain)
		}
	}

	return r0
}

// MockObserver_ObserveMerkleRoots_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObserveMerkleRoots'
type MockObserver_ObserveMerkleRoots_Call struct {
	*mock.Call
}

// ObserveMerkleRoots is a helper method to define mock.On call
//   - ctx context.Context
//   - ranges []plugintypes.ChainRange
func (_e *MockObserver_Expecter) ObserveMerkleRoots(ctx interface{}, ranges interface{}) *MockObserver_ObserveMerkleRoots_Call {
	return &MockObserver_ObserveMerkleRoots_Call{Call: _e.mock.On("ObserveMerkleRoots", ctx, ranges)}
}

func (_c *MockObserver_ObserveMerkleRoots_Call) Run(run func(ctx context.Context, ranges []plugintypes.ChainRange)) *MockObserver_ObserveMerkleRoots_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]plugintypes.ChainRange))
	})
	return _c
}

func (_c *MockObserver_ObserveMerkleRoots_Call) Return(_a0 []ccipocr3.MerkleRootChain) *MockObserver_ObserveMerkleRoots_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObserver_ObserveMerkleRoots_Call) RunAndReturn(run func(context.Context, []plugintypes.ChainRange) []ccipocr3.MerkleRootChain) *MockObserver_ObserveMerkleRoots_Call {
	_c.Call.Return(run)
	return _c
}

// ObserveOffRampNextSeqNums provides a mock function with given fields: ctx
func (_m *MockObserver) ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ObserveOffRampNextSeqNums")
	}

	var r0 []plugintypes.SeqNumChain
	if rf, ok := ret.Get(0).(func(context.Context) []plugintypes.SeqNumChain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]plugintypes.SeqNumChain)
		}
	}

	return r0
}

// MockObserver_ObserveOffRampNextSeqNums_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObserveOffRampNextSeqNums'
type MockObserver_ObserveOffRampNextSeqNums_Call struct {
	*mock.Call
}

// ObserveOffRampNextSeqNums is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockObserver_Expecter) ObserveOffRampNextSeqNums(ctx interface{}) *MockObserver_ObserveOffRampNextSeqNums_Call {
	return &MockObserver_ObserveOffRampNextSeqNums_Call{Call: _e.mock.On("ObserveOffRampNextSeqNums", ctx)}
}

func (_c *MockObserver_ObserveOffRampNextSeqNums_Call) Run(run func(ctx context.Context)) *MockObserver_ObserveOffRampNextSeqNums_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockObserver_ObserveOffRampNextSeqNums_Call) Return(_a0 []plugintypes.SeqNumChain) *MockObserver_ObserveOffRampNextSeqNums_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObserver_ObserveOffRampNextSeqNums_Call) RunAndReturn(run func(context.Context) []plugintypes.SeqNumChain) *MockObserver_ObserveOffRampNextSeqNums_Call {
	_c.Call.Return(run)
	return _c
}

// ObserveRMNRemoteCfg provides a mock function with given fields: ctx, dstChain
func (_m *MockObserver) ObserveRMNRemoteCfg(ctx context.Context, dstChain ccipocr3.ChainSelector) types.RemoteConfig {
	ret := _m.Called(ctx, dstChain)

	if len(ret) == 0 {
		panic("no return value specified for ObserveRMNRemoteCfg")
	}

	var r0 types.RemoteConfig
	if rf, ok := ret.Get(0).(func(context.Context, ccipocr3.ChainSelector) types.RemoteConfig); ok {
		r0 = rf(ctx, dstChain)
	} else {
		r0 = ret.Get(0).(types.RemoteConfig)
	}

	return r0
}

// MockObserver_ObserveRMNRemoteCfg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObserveRMNRemoteCfg'
type MockObserver_ObserveRMNRemoteCfg_Call struct {
	*mock.Call
}

// ObserveRMNRemoteCfg is a helper method to define mock.On call
//   - ctx context.Context
//   - dstChain ccipocr3.ChainSelector
func (_e *MockObserver_Expecter) ObserveRMNRemoteCfg(ctx interface{}, dstChain interface{}) *MockObserver_ObserveRMNRemoteCfg_Call {
	return &MockObserver_ObserveRMNRemoteCfg_Call{Call: _e.mock.On("ObserveRMNRemoteCfg", ctx, dstChain)}
}

func (_c *MockObserver_ObserveRMNRemoteCfg_Call) Run(run func(ctx context.Context, dstChain ccipocr3.ChainSelector)) *MockObserver_ObserveRMNRemoteCfg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockObserver_ObserveRMNRemoteCfg_Call) Return(_a0 types.RemoteConfig) *MockObserver_ObserveRMNRemoteCfg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObserver_ObserveRMNRemoteCfg_Call) RunAndReturn(run func(context.Context, ccipocr3.ChainSelector) types.RemoteConfig) *MockObserver_ObserveRMNRemoteCfg_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObserver creates a new instance of MockObserver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObserver(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObserver {
	mock := &MockObserver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
