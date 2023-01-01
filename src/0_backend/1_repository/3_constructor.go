package repository

import (
	models "go-typed-react/0_backend/0_models"

	"gorm.io/gorm"
)

func NewRepository(gormDB *gorm.DB) (rps *Repository, err error) {
	rps = &Repository{DB: gormDB}
	if err := rps.AutoMigrate(
		&models.File{}); err != nil {
		return nil, err
	}
	return rps, err
}
