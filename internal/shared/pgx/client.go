package pgx

import (
	"fmt"
	"portto/internal/shared/configx"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewClient creates a new PostgreSQL client using the provided application configuration.
func NewClient(appConfig *configx.Application) (*gorm.DB, func(), error) {
	// Initialize the PostgreSQL client with the provided configuration
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.Database.Host,
		appConfig.Database.Port,
		appConfig.Database.User,
		appConfig.Database.Password,
		appConfig.Database.Name,
	)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, func() {}, err
	}

	return client, func() {
		db, err2 := client.DB()
		if err2 != nil {
			return
		}

		if err = db.Close(); err != nil {
			return
		}
	}, nil
}
