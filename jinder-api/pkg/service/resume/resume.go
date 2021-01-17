package resume

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"Jinder/jinder-api/pkg/repository"
	"github.com/google/uuid"
)

type Service struct {
	repo repository.Resume
}

func (s Service) Get(resumeId uuid.UUID) (profile.Resume, error) {
	panic("implement me")
}

func (s Service) GetAll(userId uuid.UUID) ([]profile.Resume, error) {
	panic("implement me")
}

func (s Service) Create(resume profile.Resume) error {
	panic("implement me")
}

func (s Service) Update(resume profile.Resume) error {
	panic("implement me")
}

func (s Service) Delete(resumeId uuid.UUID) error {
	panic("implement me")
}

func NewService(repo repository.Resume) *Service {
	return &Service{repo: repo}
}
