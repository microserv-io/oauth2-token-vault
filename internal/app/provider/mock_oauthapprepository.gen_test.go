// Code generated by mockery v2.44.1. DO NOT EDIT.

package provider

import (
	context "context"

	oauthapp "github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	mock "github.com/stretchr/testify/mock"
)

// MockOAuthAppRepository is an autogenerated mock type for the OAuthAppRepository type
type MockOAuthAppRepository struct {
	mock.Mock
}

type MockOAuthAppRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOAuthAppRepository) EXPECT() *MockOAuthAppRepository_Expecter {
	return &MockOAuthAppRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, app
func (_m *MockOAuthAppRepository) Create(ctx context.Context, app *oauthapp.OAuthApp) error {
	ret := _m.Called(ctx, app)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *oauthapp.OAuthApp) error); ok {
		r0 = rf(ctx, app)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOAuthAppRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockOAuthAppRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - app *oauthapp.OAuthApp
func (_e *MockOAuthAppRepository_Expecter) Create(ctx interface{}, app interface{}) *MockOAuthAppRepository_Create_Call {
	return &MockOAuthAppRepository_Create_Call{Call: _e.mock.On("Create", ctx, app)}
}

func (_c *MockOAuthAppRepository_Create_Call) Run(run func(ctx context.Context, app *oauthapp.OAuthApp)) *MockOAuthAppRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*oauthapp.OAuthApp))
	})
	return _c
}

func (_c *MockOAuthAppRepository_Create_Call) Return(_a0 error) *MockOAuthAppRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOAuthAppRepository_Create_Call) RunAndReturn(run func(context.Context, *oauthapp.OAuthApp) error) *MockOAuthAppRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockOAuthAppRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOAuthAppRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockOAuthAppRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockOAuthAppRepository_Expecter) Delete(ctx interface{}, id interface{}) *MockOAuthAppRepository_Delete_Call {
	return &MockOAuthAppRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockOAuthAppRepository_Delete_Call) Run(run func(ctx context.Context, id string)) *MockOAuthAppRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOAuthAppRepository_Delete_Call) Return(_a0 error) *MockOAuthAppRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOAuthAppRepository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockOAuthAppRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: ctx, ownerID, id
func (_m *MockOAuthAppRepository) Find(ctx context.Context, ownerID string, id string) (*oauthapp.OAuthApp, error) {
	ret := _m.Called(ctx, ownerID, id)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *oauthapp.OAuthApp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*oauthapp.OAuthApp, error)); ok {
		return rf(ctx, ownerID, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *oauthapp.OAuthApp); ok {
		r0 = rf(ctx, ownerID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauthapp.OAuthApp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ownerID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOAuthAppRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockOAuthAppRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - ctx context.Context
//   - ownerID string
//   - id string
func (_e *MockOAuthAppRepository_Expecter) Find(ctx interface{}, ownerID interface{}, id interface{}) *MockOAuthAppRepository_Find_Call {
	return &MockOAuthAppRepository_Find_Call{Call: _e.mock.On("Find", ctx, ownerID, id)}
}

func (_c *MockOAuthAppRepository_Find_Call) Run(run func(ctx context.Context, ownerID string, id string)) *MockOAuthAppRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockOAuthAppRepository_Find_Call) Return(_a0 *oauthapp.OAuthApp, _a1 error) *MockOAuthAppRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOAuthAppRepository_Find_Call) RunAndReturn(run func(context.Context, string, string) (*oauthapp.OAuthApp, error)) *MockOAuthAppRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// ListForOwner provides a mock function with given fields: ctx, ownerID
func (_m *MockOAuthAppRepository) ListForOwner(ctx context.Context, ownerID string) ([]*oauthapp.OAuthApp, error) {
	ret := _m.Called(ctx, ownerID)

	if len(ret) == 0 {
		panic("no return value specified for ListForOwner")
	}

	var r0 []*oauthapp.OAuthApp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*oauthapp.OAuthApp, error)); ok {
		return rf(ctx, ownerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*oauthapp.OAuthApp); ok {
		r0 = rf(ctx, ownerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*oauthapp.OAuthApp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ownerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOAuthAppRepository_ListForOwner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListForOwner'
type MockOAuthAppRepository_ListForOwner_Call struct {
	*mock.Call
}

// ListForOwner is a helper method to define mock.On call
//   - ctx context.Context
//   - ownerID string
func (_e *MockOAuthAppRepository_Expecter) ListForOwner(ctx interface{}, ownerID interface{}) *MockOAuthAppRepository_ListForOwner_Call {
	return &MockOAuthAppRepository_ListForOwner_Call{Call: _e.mock.On("ListForOwner", ctx, ownerID)}
}

func (_c *MockOAuthAppRepository_ListForOwner_Call) Run(run func(ctx context.Context, ownerID string)) *MockOAuthAppRepository_ListForOwner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOAuthAppRepository_ListForOwner_Call) Return(_a0 []*oauthapp.OAuthApp, _a1 error) *MockOAuthAppRepository_ListForOwner_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOAuthAppRepository_ListForOwner_Call) RunAndReturn(run func(context.Context, string) ([]*oauthapp.OAuthApp, error)) *MockOAuthAppRepository_ListForOwner_Call {
	_c.Call.Return(run)
	return _c
}

// ListForProvider provides a mock function with given fields: ctx, providerID
func (_m *MockOAuthAppRepository) ListForProvider(ctx context.Context, providerID string) ([]*oauthapp.OAuthApp, error) {
	ret := _m.Called(ctx, providerID)

	if len(ret) == 0 {
		panic("no return value specified for ListForProvider")
	}

	var r0 []*oauthapp.OAuthApp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*oauthapp.OAuthApp, error)); ok {
		return rf(ctx, providerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*oauthapp.OAuthApp); ok {
		r0 = rf(ctx, providerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*oauthapp.OAuthApp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, providerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOAuthAppRepository_ListForProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListForProvider'
type MockOAuthAppRepository_ListForProvider_Call struct {
	*mock.Call
}

// ListForProvider is a helper method to define mock.On call
//   - ctx context.Context
//   - providerID string
func (_e *MockOAuthAppRepository_Expecter) ListForProvider(ctx interface{}, providerID interface{}) *MockOAuthAppRepository_ListForProvider_Call {
	return &MockOAuthAppRepository_ListForProvider_Call{Call: _e.mock.On("ListForProvider", ctx, providerID)}
}

func (_c *MockOAuthAppRepository_ListForProvider_Call) Run(run func(ctx context.Context, providerID string)) *MockOAuthAppRepository_ListForProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOAuthAppRepository_ListForProvider_Call) Return(_a0 []*oauthapp.OAuthApp, _a1 error) *MockOAuthAppRepository_ListForProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOAuthAppRepository_ListForProvider_Call) RunAndReturn(run func(context.Context, string) ([]*oauthapp.OAuthApp, error)) *MockOAuthAppRepository_ListForProvider_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateByID provides a mock function with given fields: ctx, id, updateFn
func (_m *MockOAuthAppRepository) UpdateByID(ctx context.Context, id string, updateFn func(*oauthapp.OAuthApp) error) error {
	ret := _m.Called(ctx, id, updateFn)

	if len(ret) == 0 {
		panic("no return value specified for UpdateByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, func(*oauthapp.OAuthApp) error) error); ok {
		r0 = rf(ctx, id, updateFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOAuthAppRepository_UpdateByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateByID'
type MockOAuthAppRepository_UpdateByID_Call struct {
	*mock.Call
}

// UpdateByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - updateFn func(*oauthapp.OAuthApp) error
func (_e *MockOAuthAppRepository_Expecter) UpdateByID(ctx interface{}, id interface{}, updateFn interface{}) *MockOAuthAppRepository_UpdateByID_Call {
	return &MockOAuthAppRepository_UpdateByID_Call{Call: _e.mock.On("UpdateByID", ctx, id, updateFn)}
}

func (_c *MockOAuthAppRepository_UpdateByID_Call) Run(run func(ctx context.Context, id string, updateFn func(*oauthapp.OAuthApp) error)) *MockOAuthAppRepository_UpdateByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(func(*oauthapp.OAuthApp) error))
	})
	return _c
}

func (_c *MockOAuthAppRepository_UpdateByID_Call) Return(_a0 error) *MockOAuthAppRepository_UpdateByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOAuthAppRepository_UpdateByID_Call) RunAndReturn(run func(context.Context, string, func(*oauthapp.OAuthApp) error) error) *MockOAuthAppRepository_UpdateByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOAuthAppRepository creates a new instance of MockOAuthAppRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOAuthAppRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOAuthAppRepository {
	mock := &MockOAuthAppRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}