// Code generated by mockery v2.43.2. DO NOT EDIT.

package plugincommon

import (
	ccipocr3 "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"
	commontypes "github.com/goplugin/plugin-libocr/commontypes"

	mapset "github.com/deckarep/golang-set/v2"

	mock "github.com/stretchr/testify/mock"
)

// MockChainSupport is an autogenerated mock type for the ChainSupport type
type MockChainSupport struct {
	mock.Mock
}

type MockChainSupport_Expecter struct {
	mock *mock.Mock
}

func (_m *MockChainSupport) EXPECT() *MockChainSupport_Expecter {
	return &MockChainSupport_Expecter{mock: &_m.Mock}
}

// KnownSourceChainsSlice provides a mock function with given fields:
func (_m *MockChainSupport) KnownSourceChainsSlice() ([]ccipocr3.ChainSelector, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for KnownSourceChainsSlice")
	}

	var r0 []ccipocr3.ChainSelector
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]ccipocr3.ChainSelector, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []ccipocr3.ChainSelector); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipocr3.ChainSelector)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainSupport_KnownSourceChainsSlice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KnownSourceChainsSlice'
type MockChainSupport_KnownSourceChainsSlice_Call struct {
	*mock.Call
}

// KnownSourceChainsSlice is a helper method to define mock.On call
func (_e *MockChainSupport_Expecter) KnownSourceChainsSlice() *MockChainSupport_KnownSourceChainsSlice_Call {
	return &MockChainSupport_KnownSourceChainsSlice_Call{Call: _e.mock.On("KnownSourceChainsSlice")}
}

func (_c *MockChainSupport_KnownSourceChainsSlice_Call) Run(run func()) *MockChainSupport_KnownSourceChainsSlice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockChainSupport_KnownSourceChainsSlice_Call) Return(_a0 []ccipocr3.ChainSelector, _a1 error) *MockChainSupport_KnownSourceChainsSlice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainSupport_KnownSourceChainsSlice_Call) RunAndReturn(run func() ([]ccipocr3.ChainSelector, error)) *MockChainSupport_KnownSourceChainsSlice_Call {
	_c.Call.Return(run)
	return _c
}

// SupportedChains provides a mock function with given fields: oracleID
func (_m *MockChainSupport) SupportedChains(oracleID commontypes.OracleID) (mapset.Set[ccipocr3.ChainSelector], error) {
	ret := _m.Called(oracleID)

	if len(ret) == 0 {
		panic("no return value specified for SupportedChains")
	}

	var r0 mapset.Set[ccipocr3.ChainSelector]
	var r1 error
	if rf, ok := ret.Get(0).(func(commontypes.OracleID) (mapset.Set[ccipocr3.ChainSelector], error)); ok {
		return rf(oracleID)
	}
	if rf, ok := ret.Get(0).(func(commontypes.OracleID) mapset.Set[ccipocr3.ChainSelector]); ok {
		r0 = rf(oracleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mapset.Set[ccipocr3.ChainSelector])
		}
	}

	if rf, ok := ret.Get(1).(func(commontypes.OracleID) error); ok {
		r1 = rf(oracleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainSupport_SupportedChains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SupportedChains'
type MockChainSupport_SupportedChains_Call struct {
	*mock.Call
}

// SupportedChains is a helper method to define mock.On call
//   - oracleID commontypes.OracleID
func (_e *MockChainSupport_Expecter) SupportedChains(oracleID interface{}) *MockChainSupport_SupportedChains_Call {
	return &MockChainSupport_SupportedChains_Call{Call: _e.mock.On("SupportedChains", oracleID)}
}

func (_c *MockChainSupport_SupportedChains_Call) Run(run func(oracleID commontypes.OracleID)) *MockChainSupport_SupportedChains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(commontypes.OracleID))
	})
	return _c
}

func (_c *MockChainSupport_SupportedChains_Call) Return(_a0 mapset.Set[ccipocr3.ChainSelector], _a1 error) *MockChainSupport_SupportedChains_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainSupport_SupportedChains_Call) RunAndReturn(run func(commontypes.OracleID) (mapset.Set[ccipocr3.ChainSelector], error)) *MockChainSupport_SupportedChains_Call {
	_c.Call.Return(run)
	return _c
}

// SupportsDestChain provides a mock function with given fields: oracle
func (_m *MockChainSupport) SupportsDestChain(oracle commontypes.OracleID) (bool, error) {
	ret := _m.Called(oracle)

	if len(ret) == 0 {
		panic("no return value specified for SupportsDestChain")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(commontypes.OracleID) (bool, error)); ok {
		return rf(oracle)
	}
	if rf, ok := ret.Get(0).(func(commontypes.OracleID) bool); ok {
		r0 = rf(oracle)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(commontypes.OracleID) error); ok {
		r1 = rf(oracle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainSupport_SupportsDestChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SupportsDestChain'
type MockChainSupport_SupportsDestChain_Call struct {
	*mock.Call
}

// SupportsDestChain is a helper method to define mock.On call
//   - oracle commontypes.OracleID
func (_e *MockChainSupport_Expecter) SupportsDestChain(oracle interface{}) *MockChainSupport_SupportsDestChain_Call {
	return &MockChainSupport_SupportsDestChain_Call{Call: _e.mock.On("SupportsDestChain", oracle)}
}

func (_c *MockChainSupport_SupportsDestChain_Call) Run(run func(oracle commontypes.OracleID)) *MockChainSupport_SupportsDestChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(commontypes.OracleID))
	})
	return _c
}

func (_c *MockChainSupport_SupportsDestChain_Call) Return(_a0 bool, _a1 error) *MockChainSupport_SupportsDestChain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainSupport_SupportsDestChain_Call) RunAndReturn(run func(commontypes.OracleID) (bool, error)) *MockChainSupport_SupportsDestChain_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockChainSupport creates a new instance of MockChainSupport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChainSupport(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChainSupport {
	mock := &MockChainSupport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
