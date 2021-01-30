package service

import (
	"Jinder/jinder-api/jobs/pkg/infrastructure"
	"Jinder/jinder-api/jobs/pkg/infrastructure/event"
	"Jinder/jinder-api/jobs/pkg/repository"
	"Jinder/jinder-api/jobs/pkg/service/dto/vacancy"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type VacancyService struct {
	repo repository.Vacancy
	infr infrastructure.Recommendation
}

func (s VacancyService) ViewVacancy(vacancyId uuid.UUID) error {
	model, err := s.repo.Get(vacancyId)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = s.infr.NotifyWhenVacancyViewed(event.NewVacancyViewed(model.UserId, model.Id))
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (s VacancyService) GetVacancy(vacancyId uuid.UUID) (vacancy.Vacancy, error) {
	model, err := s.repo.Get(vacancyId)
	if err != nil {
		log.Error(err.Error())
		return vacancy.Vacancy{}, err
	}

	return vacancy.ToDto(model), nil
}

func (s VacancyService) GetVacancies(userId uuid.UUID, count, offset string) ([]vacancy.Vacancy, error) {
	vacancies, err := s.repo.GetWithPaging(userId, count, offset)
	if err != nil {
		log.Error(err.Error())
		return []vacancy.Vacancy{}, err
	}

	var dtos []vacancy.Vacancy
	for _, v := range vacancies {
		dtos = append(dtos, vacancy.ToDto(v))
	}

	return dtos, nil
}

func (s VacancyService) CreateVacancy(vacancyDto vacancy.Vacancy) (uuid.UUID, error) {
	id, err := s.repo.Create(vacancy.ToDomain(vacancyDto))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	err = s.infr.NotifyWhenVacancyAdded(event.NewVacancyCreated(vacancyDto.UserId, id))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	return id, nil
}

func (s VacancyService) UpdateVacancy(vacancyDto vacancy.Vacancy) (uuid.UUID, error) {
	id, err := s.repo.Update(vacancy.ToDomain(vacancyDto))
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

func NewVacancyService(repo repository.Vacancy, infr infrastructure.Recommendation) *VacancyService {
	return &VacancyService{repo: repo, infr: infr}
}
