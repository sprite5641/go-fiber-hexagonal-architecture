package postgres

import (
	"go-hexagonal/internal/domain"

	"gorm.io/gorm"
)

func MigratePostgresDB(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
	)
}
