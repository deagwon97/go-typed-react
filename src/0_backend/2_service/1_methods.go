package service

import models "go-typed-react/0_backend/0_models"

func (svc *Service) CreateFile(models.File) {
	file := models.File{}
	svc.rps.CreateFile(file)
}
