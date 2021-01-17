package service

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"Jinder/jinder-api/pkg/domain/registration"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service/resume"
	"Jinder/jinder-api/pkg/service/token_manager"
	"Jinder/jinder-api/pkg/service/user"
	"Jinder/jinder-api/pkg/service/vacancy"
	"github.com/google/uuid"
)

type User interface {
	Register(registration.User) error
	Update(registration.User) error
	Get(userId uuid.UUID) (registration.User, error)
}

type TokenManager interface {
	Generate() (string, error)
	Validate(token string) (bool, error)
	Parse(token string) (uuid.UUID, error)
}

type Resume interface {
	Get(resumeId uuid.UUID) (profile.Resume, error)
	GetAll(userId uuid.UUID) ([]profile.Resume, error)
	Create(resume profile.Resume) error
	Update(resume profile.Resume) error
	Delete(resumeId uuid.UUID) error
}

type Vacancy interface {
	Get(vacancyId uuid.UUID) (profile.Vacancy, error)
	GetWithPaging(userId uuid.UUID) ([]profile.Vacancy, error)
	Create(vacancy profile.Vacancy) error
	Update(vacancy profile.Vacancy) error
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
