package repository

import (
	models "go-typed-react/0_backend/0_models"
)

func (rps *Repository) CreateFile(
	file models.File,
) (
	models.File, error,
) {
	err := rps.Create(&file).Error
	return file, err
}
