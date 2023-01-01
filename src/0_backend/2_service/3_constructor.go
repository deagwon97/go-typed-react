package service

import repository "go-typed-react/0_backend/1_repository"

// HandlerInterface의 생성자
func NewService(rps *repository.Repository) (
	*Service, error) {
	return &Service{
		rps: rps,
	}, nil
}
