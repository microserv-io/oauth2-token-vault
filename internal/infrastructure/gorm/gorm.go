package gorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(connectionString string, autoMigrate bool) (*gorm.DB, error) {

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
