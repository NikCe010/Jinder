package mocks

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	"Jinder/jinder-api/jobs/pkg/tests/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockVacancyRepository struct {
	mock.Mock
}

func (m MockVacancyRepository) Get(vacancyId uuid.UUID) (profile.Vacancy, error) {
	return models.Vacancy, nil
}

func (m MockVacancyRepository) GetWithPaging(userId uuid.UUID, count string, page string) ([]profile.Vacancy, error) {
	var vacancies []profile.Vacancy
	vacancies = append(vacancies, models.Vacancy)
	return vacancies, nil
}

func (m MockVacancyRepository) Create(vacancy profile.Vacancy) (uuid.UUID, error) {
	return vacancy.Id, nil
}

func (m MockVacancyRepository) Update(vacancy profile.Vacancy) (uuid.UUID, error) {
	return vacancy.Id, nil
}

func (m MockVacancyRepository) Delete(vacancyId uuid.UUID) error {
	return nil
}
