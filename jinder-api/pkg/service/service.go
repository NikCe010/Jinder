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
	Update(dto.User) (uuid.UUID, error)
	Get(userId uuid.UUID) (dto.User, error)
}

type TokenManager interface {
	Generate() (string, error)
	Validate(token string) (bool, error)
	Parse(token string) (uuid.UUID, error)
}

type Resume interface {
	Get(resumeId uuid.UUID) (dto.Resume, error)
	GetAll(userId uuid.UUID) ([]dto.Resume, error)
	Create(resume dto.Resume) (uuid.UUID, error)
	Update(resume dto.Resume) (uuid.UUID, error)
	Delete(resumeId uuid.UUID) error
}

type Vacancy interface {
	Get(vacancyId uuid.UUID) (dto.Vacancy, error)
	GetWithPaging(userId uuid.UUID, count int, offset int) ([]dto.Vacancy, error)
	Create(vacancy dto.Vacancy) (uuid.UUID, error)
	Update(vacancy dto.Vacancy) (uuid.UUID, error)
	Delete(vacancyId uuid.UUID) error
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
