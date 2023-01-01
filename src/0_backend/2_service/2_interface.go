package service

import models "go-typed-react/0_backend/0_models"

type ServiceInterface interface {
	CreateFile(models.File)
}
