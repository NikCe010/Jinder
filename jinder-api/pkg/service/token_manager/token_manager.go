package token_manager

import "github.com/google/uuid"

type TokenManagerService struct {
}

func NewService() *TokenManagerService {
	return &TokenManagerService{}
}

func (s TokenManagerService) Generate() (string, error) {
	panic("implement me")
}

func (s TokenManagerService) Validate(token string) (bool, error) {
	panic("implement me")
}

func (s TokenManagerService) Parse(token string) (uuid.UUID, error) {
	panic("implement me")
}
