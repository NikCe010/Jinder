package service

import (
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service/dto"
	"Jinder/jinder-api/pkg/service/resume"
	"Jinder/jinder-api/pkg/service/token_manager"
	"Jinder/jinder-api/pkg/service/user"
	"Jinder/jinder-api/pkg/service/vacancy"
	"github.com/google/uuid"
)

type User interface {
	Register(dto.User) (uuid.UUID, error)
	UpdateUser(dto.User) (uuid.UUID, error)
	GetUser(userId uuid.UUID) (dto.User, error)
}

type TokenManager interface {
	Generate() (string, error)
	Validate(token string) (bool, error)
	Parse(token string) (uuid.UUID, error)
}

type Resume interface {
	GetResume(resumeId uuid.UUID) (dto.Resume, error)
	GetResumes(userId uuid.UUID, count int, offset int) ([]dto.Resume, error)
	CreateResume(resume dto.Resume) (uuid.UUID, error)
	UpdateResume(resume dto.Resume) (uuid.UUID, error)
	DeleteResume(resumeId uuid.UUID) error
}

type Vacancy interface {
	GetVacancy(vacancyId uuid.UUID) (dto.Vacancy, error)
	GetVacancies(userId uuid.UUID, count int, offset int) ([]dto.Vacancy, error)
	CreateVacancy(vacancy dto.Vacancy) (uuid.UUID, error)
	UpdateVacancy(vacancy dto.Vacancy) (uuid.UUID, error)
	DeleteVacancy(vacancyId uuid.UUID) error
}

type Service struct {
	User
	Resume
	Vacancy
	TokenManager
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:         user.NewService(repos.User),
		Resume:       resume.NewService(repos.Resume),
		Vacancy:      vacancy.NewService(repos.Vacancy),
		TokenManager: token_manager.NewService(),
	}
}
