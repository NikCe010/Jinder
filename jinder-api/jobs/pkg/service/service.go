package service

import (
	"Jinder/jinder-api/jobs/pkg/infrastructure"
	"Jinder/jinder-api/jobs/pkg/repository"
	"Jinder/jinder-api/jobs/pkg/service/dto/resume"
	"Jinder/jinder-api/jobs/pkg/service/dto/user"
	"Jinder/jinder-api/jobs/pkg/service/dto/vacancy"
	"github.com/google/uuid"
)

type User interface {
	Register(user.User) (uuid.UUID, error)
	UpdateUser(user.User) (uuid.UUID, error)
	GetUser(userId uuid.UUID) (user.User, error)
}

type TokenManager interface {
	Generate(email, password string) (string, error)
	Validate(token string) (uuid.UUID, error)
}

type Resume interface {
	GetResume(resumeId uuid.UUID) (resume.Resume, error)
	GetResumes(userId uuid.UUID, count, offset string) ([]resume.Resume, error)
	CreateResume(resume resume.Resume) (uuid.UUID, error)
	UpdateResume(resume resume.Resume) (uuid.UUID, error)
	DeleteResume(resumeId uuid.UUID) error
	ViewResume(resumeId uuid.UUID) error
}

type Vacancy interface {
	GetVacancy(vacancyId uuid.UUID) (vacancy.Vacancy, error)
	GetVacancies(userId uuid.UUID, count, offset string) ([]vacancy.Vacancy, error)
	CreateVacancy(vacancy vacancy.Vacancy) (uuid.UUID, error)
	UpdateVacancy(vacancy vacancy.Vacancy) (uuid.UUID, error)
	DeleteVacancy(vacancyId uuid.UUID) error
	ViewVacancy(vacancyId uuid.UUID) error
}

type Service struct {
	User
	Resume
	Vacancy
	TokenManager
}

func NewService(repos *repository.Repository, infrastructure *infrastructure.Infrastructure) *Service {
	return &Service{
		User:         NewUserService(repos.User),
		Resume:       NewResumeService(repos.Resume, infrastructure),
		Vacancy:      NewVacancyService(repos.Vacancy, infrastructure),
		TokenManager: NewTokenManagerService(repos.User),
	}
}
