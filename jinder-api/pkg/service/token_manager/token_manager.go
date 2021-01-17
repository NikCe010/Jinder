package token_manager

import "github.com/google/uuid"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) Generate() (string, error) {
	panic("implement me")
}

func (s Service) Validate(token string) (bool, error) {
	panic("implement me")
}

func (s Service) Parse(token string) (uuid.UUID, error) {
	panic("implement me")
}
