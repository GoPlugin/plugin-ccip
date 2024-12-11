// Code generated by mockery v2.43.2. DO NOT EDIT.

package contractreader

import (
	context "context"

	contractreader "github.com/goplugin/plugin-ccip/pkg/contractreader"
	mock "github.com/stretchr/testify/mock"

	primitives "github.com/goplugin/plugin-common/pkg/types/query/primitives"

	query "github.com/goplugin/plugin-common/pkg/types/query"

	types "github.com/goplugin/plugin-common/pkg/types"
)

// MockExtended is an autogenerated mock type for the Extended type
type MockExtended struct {
	mock.Mock
}

type MockExtended_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExtended) EXPECT() *MockExtended_Expecter {
	return &MockExtended_Expecter{mock: &_m.Mock}
}

// Bind provides a mock function with given fields: ctx, bindings
func (_m *MockExtended) Bind(ctx context.Context, bindings []types.BoundContract) error {
	ret := _m.Called(ctx, bindings)

	if len(ret) == 0 {
		panic("no return value specified for Bind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.BoundContract) error); ok {
		r0 = rf(ctx, bindings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExtended_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type MockExtended_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
//   - ctx context.Context
//   - bindings []types.BoundContract
func (_e *MockExtended_Expecter) Bind(ctx interface{}, bindings interface{}) *MockExtended_Bind_Call {
	return &MockExtended_Bind_Call{Call: _e.mock.On("Bind", ctx, bindings)}
}

func (_c *MockExtended_Bind_Call) Run(run func(ctx context.Context, bindings []types.BoundContract)) *MockExtended_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BoundContract))
	})
	return _c
}

func (_c *MockExtended_Bind_Call) Return(_a0 error) *MockExtended_Bind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Bind_Call) RunAndReturn(run func(context.Context, []types.BoundContract) error) *MockExtended_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// ExtendedGetLatestValue provides a mock function with given fields: ctx, contractName, methodName, confidenceLevel, params, returnVal
func (_m *MockExtended) ExtendedGetLatestValue(ctx context.Context, contractName string, methodName string, confidenceLevel primitives.ConfidenceLevel, params interface{}, returnVal interface{}) error {
	ret := _m.Called(ctx, contractName, methodName, confidenceLevel, params, returnVal)

	if len(ret) == 0 {
		panic("no return value specified for ExtendedGetLatestValue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, primitives.ConfidenceLevel, interface{}, interface{}) error); ok {
		r0 = rf(ctx, contractName, methodName, confidenceLevel, params, returnVal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExtended_ExtendedGetLatestValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExtendedGetLatestValue'
type MockExtended_ExtendedGetLatestValue_Call struct {
	*mock.Call
}

// ExtendedGetLatestValue is a helper method to define mock.On call
//   - ctx context.Context
//   - contractName string
//   - methodName string
//   - confidenceLevel primitives.ConfidenceLevel
//   - params interface{}
//   - returnVal interface{}
func (_e *MockExtended_Expecter) ExtendedGetLatestValue(ctx interface{}, contractName interface{}, methodName interface{}, confidenceLevel interface{}, params interface{}, returnVal interface{}) *MockExtended_ExtendedGetLatestValue_Call {
	return &MockExtended_ExtendedGetLatestValue_Call{Call: _e.mock.On("ExtendedGetLatestValue", ctx, contractName, methodName, confidenceLevel, params, returnVal)}
}

func (_c *MockExtended_ExtendedGetLatestValue_Call) Run(run func(ctx context.Context, contractName string, methodName string, confidenceLevel primitives.ConfidenceLevel, params interface{}, returnVal interface{})) *MockExtended_ExtendedGetLatestValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(primitives.ConfidenceLevel), args[4].(interface{}), args[5].(interface{}))
	})
	return _c
}

func (_c *MockExtended_ExtendedGetLatestValue_Call) Return(_a0 error) *MockExtended_ExtendedGetLatestValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_ExtendedGetLatestValue_Call) RunAndReturn(run func(context.Context, string, string, primitives.ConfidenceLevel, interface{}, interface{}) error) *MockExtended_ExtendedGetLatestValue_Call {
	_c.Call.Return(run)
	return _c
}

// ExtendedQueryKey provides a mock function with given fields: ctx, contractName, filter, limitAndSort, sequenceDataType
func (_m *MockExtended) ExtendedQueryKey(ctx context.Context, contractName string, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{}) ([]types.Sequence, error) {
	ret := _m.Called(ctx, contractName, filter, limitAndSort, sequenceDataType)

	if len(ret) == 0 {
		panic("no return value specified for ExtendedQueryKey")
	}

	var r0 []types.Sequence
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) ([]types.Sequence, error)); ok {
		return rf(ctx, contractName, filter, limitAndSort, sequenceDataType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) []types.Sequence); ok {
		r0 = rf(ctx, contractName, filter, limitAndSort, sequenceDataType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Sequence)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) error); ok {
		r1 = rf(ctx, contractName, filter, limitAndSort, sequenceDataType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExtended_ExtendedQueryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExtendedQueryKey'
type MockExtended_ExtendedQueryKey_Call struct {
	*mock.Call
}

// ExtendedQueryKey is a helper method to define mock.On call
//   - ctx context.Context
//   - contractName string
//   - filter query.KeyFilter
//   - limitAndSort query.LimitAndSort
//   - sequenceDataType interface{}
func (_e *MockExtended_Expecter) ExtendedQueryKey(ctx interface{}, contractName interface{}, filter interface{}, limitAndSort interface{}, sequenceDataType interface{}) *MockExtended_ExtendedQueryKey_Call {
	return &MockExtended_ExtendedQueryKey_Call{Call: _e.mock.On("ExtendedQueryKey", ctx, contractName, filter, limitAndSort, sequenceDataType)}
}

func (_c *MockExtended_ExtendedQueryKey_Call) Run(run func(ctx context.Context, contractName string, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{})) *MockExtended_ExtendedQueryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(query.KeyFilter), args[3].(query.LimitAndSort), args[4].(interface{}))
	})
	return _c
}

func (_c *MockExtended_ExtendedQueryKey_Call) Return(_a0 []types.Sequence, _a1 error) *MockExtended_ExtendedQueryKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtended_ExtendedQueryKey_Call) RunAndReturn(run func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) ([]types.Sequence, error)) *MockExtended_ExtendedQueryKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetBindings provides a mock function with given fields: contractName
func (_m *MockExtended) GetBindings(contractName string) []contractreader.ExtendedBoundContract {
	ret := _m.Called(contractName)

	if len(ret) == 0 {
		panic("no return value specified for GetBindings")
	}

	var r0 []contractreader.ExtendedBoundContract
	if rf, ok := ret.Get(0).(func(string) []contractreader.ExtendedBoundContract); ok {
		r0 = rf(contractName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contractreader.ExtendedBoundContract)
		}
	}

	return r0
}

// MockExtended_GetBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBindings'
type MockExtended_GetBindings_Call struct {
	*mock.Call
}

// GetBindings is a helper method to define mock.On call
//   - contractName string
func (_e *MockExtended_Expecter) GetBindings(contractName interface{}) *MockExtended_GetBindings_Call {
	return &MockExtended_GetBindings_Call{Call: _e.mock.On("GetBindings", contractName)}
}

func (_c *MockExtended_GetBindings_Call) Run(run func(contractName string)) *MockExtended_GetBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockExtended_GetBindings_Call) Return(_a0 []contractreader.ExtendedBoundContract) *MockExtended_GetBindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_GetBindings_Call) RunAndReturn(run func(string) []contractreader.ExtendedBoundContract) *MockExtended_GetBindings_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExtended creates a new instance of MockExtended. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExtended(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExtended {
	mock := &MockExtended{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}