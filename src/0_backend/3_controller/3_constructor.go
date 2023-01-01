package controller

import service "go-typed-react/0_backend/2_service"

func NewController(svc *service.Service) (*Controller, error) {
	return &Controller{
		svc: svc,
	}, nil
}
