package repository

import (
	models "go-typed-react/0_backend/0_models"
)

type RepositoryInterface interface {
	CreateFile(models.File) (models.File, error)
}
