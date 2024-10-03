package gorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	maxRetries    = 6
	retryInterval = 5 * time.Second
)

func Open(connectionString string, autoMigrate bool) (*gorm.DB, error) {
	var db *gorm.DB
	for i := 0; i < maxRetries; i++ {
		dbVar, err := gorm.Open(
			postgres.New(postgres.Config{
				DSN: connectionString,
			}),
			&gorm.Config{},
		)
		if err == nil {
			db = dbVar
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}
	if db == nil {
		log.Fatalf("Could not connect to the database after %d attempts.", maxRetries)
	}

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: connectionString,
		}),
		&gorm.Config{},
	)

	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if !autoMigrate {
		return db, nil
	}

	if err := db.AutoMigrate(&OAuthApp{}, &Provider{}); err != nil {
		return nil, fmt.Errorf("error migrating database schema: %w", err)
	}

	return db, nil
}
