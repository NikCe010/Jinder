package vacancy

import (
	"Jinder/jinder-api/pkg/domain/profile"
	domain "Jinder/jinder-api/pkg/domain/profile/shared"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service/dto"
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
)

type VacancyService struct {
	repo repository.Vacancy
}

func (s VacancyService) Get(vacancyId uuid.UUID) (dto.Vacancy, error) {
	vacancy, err := s.repo.Get(vacancyId)
	if err != nil {
		return dto.Vacancy{}, err
	}

	return MappingToDto(vacancy), nil
}

func (s VacancyService) GetWithPaging(userId uuid.UUID, count int, offset int) ([]dto.Vacancy, error) {
	vacancies, err := s.repo.GetWithPaging(userId, count, offset)
	if err != nil {
		return []dto.Vacancy{}, err
	}

	var dtos []dto.Vacancy
	for _, v := range vacancies {
		dtos = append(dtos, MappingToDto(v))
	}

	return dtos, nil
}

func (s VacancyService) Create(vacancy dto.Vacancy) (uuid.UUID, error) {
	id, err := s.repo.Create(Mapping(vacancy))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s VacancyService) Update(vacancy dto.Vacancy) (uuid.UUID, error) {
	id, err := s.repo.Update(Mapping(vacancy))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s VacancyService) Delete(vacancyId uuid.UUID) error {
	err := s.repo.Delete(vacancyId)
	if err != nil {
		return err
	}

	return nil
}

func NewService(repo repository.Vacancy) *VacancyService {
	return &VacancyService{repo: repo}
}

func Mapping(vacancy dto.Vacancy) profile.Vacancy {
	return profile.Vacancy{
		Id:                 vacancy.Id,
		UserId:             vacancy.UserId,
		ProgrammerLanguage: domain.ProgrammerLanguage(vacancy.ProgrammerLanguage),
		ProgrammerLevel:    domain.ProgrammerLevel(vacancy.ProgrammerLanguage),
		ProgrammerType:     domain.ProgrammerType(vacancy.ProgrammerLanguage),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		OtherBenefits:      vacancy.OtherBenefits,
	}
}

func MappingToDto(vacancy profile.Vacancy) dto.Vacancy {
	return dto.Vacancy{
		Id:                 vacancy.Id,
		UserId:             vacancy.UserId,
		ProgrammerLanguage: shared.ProgrammerLanguage(vacancy.ProgrammerLanguage),
		ProgrammerLevel:    shared.ProgrammerLevel(vacancy.ProgrammerLanguage),
		ProgrammerType:     shared.ProgrammerType(vacancy.ProgrammerLanguage),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		OtherBenefits:      vacancy.OtherBenefits,
	}
}
