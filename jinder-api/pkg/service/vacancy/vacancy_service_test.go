package vacancy

import (
	"Jinder/jinder-api/pkg/domain/profile"
	domain "Jinder/jinder-api/pkg/domain/profile/shared"
	"Jinder/jinder-api/pkg/service/dto"
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var vacancyId = uuid.New()

var vacancyDto = dto.Vacancy{
	Id:                 vacancyId,
	UserId:             uuid.New(),
	ProgrammerLevel:    shared.Middle,
	ProgrammerType:     shared.Backend,
	ProgrammerLanguage: shared.Golang,
	CompanyName:        "Test Company",
	SalaryFrom:         "150000",
	SalaryTo:           "200000",
	OtherBenefits:      "Medical Insurance, paid vacation 31 days",
}

var vacancy = profile.Vacancy{
	Id:                 vacancyId,
	UserId:             uuid.New(),
	ProgrammerLevel:    domain.Middle,
	ProgrammerType:     domain.Backend,
	ProgrammerLanguage: domain.Golang,
	CompanyName:        "Test Company",
	SalaryFrom:         "150000",
	SalaryTo:           "200000",
	OtherBenefits:      "Medical Insurance, paid vacation 31 days",
}

type MockVacancyRepository struct {
	mock.Mock
}

func (m MockVacancyRepository) Get(vacancyId uuid.UUID) (profile.Vacancy, error) {
	return vacancy, nil
}

func (m MockVacancyRepository) GetWithPaging(userId uuid.UUID, count int, page int) ([]profile.Vacancy, error) {
	var vacancies []profile.Vacancy
	vacancies = append(vacancies, vacancy)
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

func TestVacancyService_Create_ShouldCompleteSuccessful(t *testing.T) {
	service := NewService(MockVacancyRepository{})

	id, err := service.Create(vacancyDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestVacancyService_Update_ShouldCompleteSuccessful(t *testing.T) {
	service := NewService(MockVacancyRepository{})

	id, err := service.Update(vacancyDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestVacancyService_Get_ShouldReturnNotEmptyVacancy(t *testing.T) {
	service := NewService(MockVacancyRepository{})

	result, err := service.Get(vacancyDto.Id)

	assert.NoError(t, err)
	assert.NotEqual(t, result, profile.Vacancy{})
}

func TestVacancyService_GetWithPaging_ShouldReturnOneVacancy(t *testing.T) {
	service := NewService(MockVacancyRepository{})

	results, err := service.GetWithPaging(vacancyDto.Id, 5, 0)

	assert.NoError(t, err)
	assert.NotEqual(t, results[0], uuid.UUID{})
	assert.Equal(t, len(results), 1)
}

func TestVacancyService_Delete_ShouldCompleteSuccessful(t *testing.T) {
	service := NewService(MockVacancyRepository{})

	err := service.Delete(vacancyDto.Id)

	assert.NoError(t, err)
}
