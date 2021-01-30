package tests

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	"Jinder/jinder-api/jobs/pkg/service"
	"Jinder/jinder-api/jobs/pkg/tests/mocks"
	"Jinder/jinder-api/jobs/pkg/tests/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVacancyService_Create_ShouldCompleteSuccessful(t *testing.T) {
	service := service.NewVacancyService(mocks.MockVacancyRepository{}, mocks.MockRecommendationClient{})

	id, err := service.CreateVacancy(models.VacancyDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestVacancyService_Update_ShouldCompleteSuccessful(t *testing.T) {
	service := service.NewVacancyService(mocks.MockVacancyRepository{}, mocks.MockRecommendationClient{})

	id, err := service.UpdateVacancy(models.VacancyDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestVacancyService_Get_ShouldReturnNotEmptyVacancy(t *testing.T) {
	service := service.NewVacancyService(mocks.MockVacancyRepository{}, mocks.MockRecommendationClient{})

	result, err := service.GetVacancy(models.VacancyDto.Id)

	assert.NoError(t, err)
	assert.NotEqual(t, result, profile.Vacancy{})
}

func TestVacancyService_GetWithPaging_ShouldReturnOneVacancy(t *testing.T) {
	service := service.NewVacancyService(mocks.MockVacancyRepository{}, mocks.MockRecommendationClient{})

	results, err := service.GetVacancies(models.VacancyDto.Id, "5", "0")

	assert.NoError(t, err)
	assert.NotEqual(t, results[0], uuid.UUID{})
	assert.Equal(t, len(results), 1)
}

func TestVacancyService_Delete_ShouldCompleteSuccessful(t *testing.T) {
	service := service.NewVacancyService(mocks.MockVacancyRepository{}, mocks.MockRecommendationClient{})

	err := service.DeleteVacancy(models.VacancyDto.Id)

	assert.NoError(t, err)
}
