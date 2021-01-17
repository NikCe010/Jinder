package vacancy

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"Jinder/jinder-api/pkg/repository"
	"github.com/google/uuid"
)

type Service struct {
	repo repository.Vacancy
}

func (s Service) Get(vacancyId uuid.UUID) (profile.Vacancy, error) {
	panic("implement me")
}

func (s Service) GetWithPaging(userId uuid.UUID) ([]profile.Vacancy, error) {
	panic("implement me")
}

func (s Service) Create(vacancy profile.Vacancy) error {
	panic("implement me")
}

func (s Service) Update(vacancy profile.Vacancy) error {
	panic("implement me")
}

func (s Service) Delete(vacancyId uuid.UUID) error {
	panic("implement me")
}

func NewService(repo repository.Vacancy) *Service {
	return &Service{repo: repo}
}
