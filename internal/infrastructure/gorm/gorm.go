package gorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	defaultRetryInterval = 5 * time.Second
)

// Open opens a connection to the database and returns a gorm.DB instance.
func Open(
	connectionString string,
	autoMigrate bool,
	maxAttempts int,
	retryInterval time.Duration,
	openFunc func(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error),
) (*gorm.DB, error) {

	if connectionString == "" {
		return nil, fmt.Errorf("connection string is required")
	}

	if maxAttempts == 0 {
		return nil, fmt.Errorf("max attempts must be greater than 0")
	}

	if retryInterval == 0 {
		retryInterval = defaultRetryInterval
	}

	if openFunc == nil {
		openFunc = gorm.Open
	}

	var db *gorm.DB
	i := 0
	for {
		dbVar, err := openFunc(
			postgres.New(postgres.Config{
				DSN: connectionString,
			}),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
			},
		)
		if err == nil {
			db = dbVar
			break
		}

		i++

		if i == maxAttempts {
			break
		}

		time.Sleep(retryInterval * time.Second)
	}
	if db == nil {
		return nil, fmt.Errorf("could not connect to the database after %d attempts", maxAttempts)
	}

	if !autoMigrate {
		return db, nil
	}

	if err := db.AutoMigrate(&OAuthApp{}, &Provider{}); err != nil {
		return nil, fmt.Errorf("error migrating database schema: %w", err)
	}

	return db, nil
}
