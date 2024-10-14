package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

type GRPCScenario struct {
	Name           string
	BeforeTestFunc func(t *testing.T, app *TestApp)
	Input          interface{}
	Request        func(conn *grpc.ClientConn, input interface{}) (interface{}, error)
	ExpectedResp   interface{}
	ExpectedErr    string
}

func (tc *GRPCScenario) Test(t *testing.T) {
	t.Run(tc.Name, tc.test)
}

func (tc *GRPCScenario) test(t *testing.T) {
	app := NewTestApp("config")
	defer app.Cleanup()

	if tc.BeforeTestFunc != nil {
		tc.BeforeTestFunc(t, app)
	}

	// Set up the gRPC connection
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", app.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	response, err := tc.Request(conn, tc.Input)

	if tc.ExpectedErr == "" {
		assert.Nil(t, err)
	} else {
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), tc.ExpectedErr)
	}

	if tc.ExpectedResp != nil {
		assert.EqualExportedValues(t, tc.ExpectedResp, response)
	} else {
		assert.Nil(t, response)
	}
}
