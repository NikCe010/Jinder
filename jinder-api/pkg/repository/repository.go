package repository

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"Jinder/jinder-api/pkg/domain/registration"
	"Jinder/jinder-api/pkg/repository/resume"
	"Jinder/jinder-api/pkg/repository/user"
	"Jinder/jinder-api/pkg/repository/vacancy"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User interface {
	//Register new user.
	//If failed return error.
	Register(registration.User) error

	//Update user.
	//If failed return error.
	Update(registration.User) error

	//Get user by id.
	//Return user and error.
	Get(userId uuid.UUID) (registration.User, error)
}

type Resume interface {
	//Get resume by resume id.
	//Return resume and error.
	Get(resumeId uuid.UUID) (profile.Resume, error)

	//Get all resumes by user id.
	//Return slice of resumes (max 10) and error.
	GetAll(userId uuid.UUID) ([]profile.Resume, error)

	//Create resume.
	//If failed return error.
	Create(resume profile.Resume) (uuid.UUID, error)

	//Update resume.
	//If failed return error.
	Update(resume profile.Resume) (uuid.UUID, error)

	//Delete resume by resume id.
	//If failed return error.
	Delete(resumeId uuid.UUID) error
}

type Vacancy interface {
	//Get vacancy by vacancy id.
	//Return vacancy and error.
	Get(vacancyId uuid.UUID) (profile.Vacancy, error)

	//Get vacancy by user id, page number and count.
	//Return vacancy and error.
	GetWithPaging(userId uuid.UUID, count int, page int) ([]profile.Vacancy, error)

	//Create vacancy.
	//If failed return error.
	Create(vacancy profile.Vacancy) error

	//Update vacancy.
	//If failed return error.
	Update(vacancy profile.Vacancy) error

	//Delete vacancy by vacancy id.
	//If failed return error.
	Delete(vacancyId uuid.UUID) error
}

type Repository struct {
	User
	Resume
	Vacancy
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    user.NewPostgres(db),
		Resume:  repository.NewResumePostgres(db),
		Vacancy: vacancy.NewPostgres(db),
	}
}
