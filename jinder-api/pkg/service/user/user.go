package user

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"Jinder/jinder-api/pkg/repository"
	"github.com/google/uuid"
)

type Service struct {
	repo repository.User
}

func (s Service) Register(user registration.User) error {
	panic("implement me")
}

func (s Service) Update(user registration.User) error {
	panic("implement me")
}

func (s Service) Get(userId uuid.UUID) (registration.User, error) {
	panic("implement me")
}

func NewService(repo repository.User) *Service {
	return &Service{repo: repo}
}
