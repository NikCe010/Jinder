package vacancy

import (
	"Jinder/jinder-api/pkg/domain/profile"
	domain "Jinder/jinder-api/pkg/domain/profile/shared"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service/dto"
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type VacancyService struct {
	repo repository.Vacancy
}

func (s VacancyService) GetVacancy(vacancyId uuid.UUID) (dto.Vacancy, error) {
	vacancy, err := s.repo.Get(vacancyId)
	if err != nil {
		log.Error(err.Error())
		return dto.Vacancy{}, err
	}

	return MappingToDto(vacancy), nil
}

func (s VacancyService) GetVacancies(userId uuid.UUID, count int, offset int) ([]dto.Vacancy, error) {
	vacancies, err := s.repo.GetWithPaging(userId, count, offset)
	if err != nil {
		log.Error(err.Error())
		return []dto.Vacancy{}, err
	}

	var dtos []dto.Vacancy
	for _, v := range vacancies {
		dtos = append(dtos, MappingToDto(v))
	}

	return dtos, nil
}

func (s VacancyService) CreateVacancy(vacancy dto.Vacancy) (uuid.UUID, error) {
	id, err := s.repo.Create(Mapping(vacancy))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s VacancyService) UpdateVacancy(vacancy dto.Vacancy) (uuid.UUID, error) {
	id, err := s.repo.Update(Mapping(vacancy))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s VacancyService) DeleteVacancy(vacancyId uuid.UUID) error {
	err := s.repo.Delete(vacancyId)
	if err != nil {
		log.Error(err.Error())
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
		ProgrammerLevel:    domain.ProgrammerLevel(vacancy.ProgrammerLevel),
		ProgrammerType:     domain.ProgrammerType(vacancy.ProgrammerType),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		ExtraBenefits:      vacancy.ExtraBenefits,
	}
}

func MappingToDto(vacancy profile.Vacancy) dto.Vacancy {
	return dto.Vacancy{
		Id:                 vacancy.Id,
		UserId:             vacancy.UserId,
		ProgrammerLanguage: shared.ProgrammerLanguage(vacancy.ProgrammerLanguage),
		ProgrammerLevel:    shared.ProgrammerLevel(vacancy.ProgrammerLevel),
		ProgrammerType:     shared.ProgrammerType(vacancy.ProgrammerType),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		ExtraBenefits:      vacancy.ExtraBenefits,
	}
}
