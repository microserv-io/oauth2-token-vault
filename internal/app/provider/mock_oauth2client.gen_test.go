// Code generated by mockery v2.44.1. DO NOT EDIT.

package provider

import (
	context "context"

	oauth2 "github.com/microserv-io/oauth-credentials-server/internal/domain/oauth2"
	mock "github.com/stretchr/testify/mock"
)

// MockOAuth2Client is an autogenerated mock type for the OAuth2Client type
type MockOAuth2Client struct {
	mock.Mock
}

type MockOAuth2Client_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOAuth2Client) EXPECT() *MockOAuth2Client_Expecter {
	return &MockOAuth2Client_Expecter{mock: &_m.Mock}
}

// Exchange provides a mock function with given fields: ctx, config, code
func (_m *MockOAuth2Client) Exchange(ctx context.Context, config *oauth2.Config, code string) (*oauth2.Token, error) {
	ret := _m.Called(ctx, config, code)

	if len(ret) == 0 {
		panic("no return value specified for Exchange")
	}

	var r0 *oauth2.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2.Config, string) (*oauth2.Token, error)); ok {
		return rf(ctx, config, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2.Config, string) *oauth2.Token); ok {
		r0 = rf(ctx, config, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oauth2.Config, string) error); ok {
		r1 = rf(ctx, config, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOAuth2Client_Exchange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exchange'
type MockOAuth2Client_Exchange_Call struct {
	*mock.Call
}

// Exchange is a helper method to define mock.On call
//   - ctx context.Context
//   - config *oauth2.Config
//   - code string
func (_e *MockOAuth2Client_Expecter) Exchange(ctx interface{}, config interface{}, code interface{}) *MockOAuth2Client_Exchange_Call {
	return &MockOAuth2Client_Exchange_Call{Call: _e.mock.On("Exchange", ctx, config, code)}
}

func (_c *MockOAuth2Client_Exchange_Call) Run(run func(ctx context.Context, config *oauth2.Config, code string)) *MockOAuth2Client_Exchange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*oauth2.Config), args[2].(string))
	})
	return _c
}

func (_c *MockOAuth2Client_Exchange_Call) Return(_a0 *oauth2.Token, _a1 error) *MockOAuth2Client_Exchange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOAuth2Client_Exchange_Call) RunAndReturn(run func(context.Context, *oauth2.Config, string) (*oauth2.Token, error)) *MockOAuth2Client_Exchange_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOAuth2Client creates a new instance of MockOAuth2Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOAuth2Client(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOAuth2Client {
	mock := &MockOAuth2Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}