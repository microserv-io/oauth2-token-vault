package gorm

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// MockGorm is a mock implementation of the gorm.DB
type MockGorm struct {
	mock.Mock
}

func (m *MockGorm) Open(dialector gorm.Dialector, opt ...gorm.Option) (*gorm.DB, error) {
	args := m.Called(dialector, opt)
	return args.Get(0).(*gorm.DB), args.Error(1)
}

func TestOpen(t *testing.T) {
	tests := []struct {
		name             string
		connectionString string
		mockReturnDB     *gorm.DB
		mockReturnErr    error
		expectedErr      bool
	}{
		{
			name:             "Success",
			connectionString: "valid_connection_string",
			mockReturnDB:     new(gorm.DB),
			mockReturnErr:    nil,
			expectedErr:      false,
		},
		{
			name:             "Failure",
			connectionString: "invalid_connection_string",
			mockReturnDB:     nil,
			mockReturnErr:    errors.New("connection error"),
			expectedErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGorm := new(MockGorm)
			mockGorm.On("Open", mock.Anything, mock.Anything).Return(tt.mockReturnDB, tt.mockReturnErr)

			openFunc := func(dialector gorm.Dialector, opt ...gorm.Option) (*gorm.DB, error) {
				return mockGorm.Open(dialector)
			}

			db, err := Open(tt.connectionString, false, 1, 0, openFunc)
			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, db)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db)
			}
			mockGorm.AssertNumberOfCalls(t, "Open", 1)
		})
	}
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Provider{}, &OAuthApp{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
