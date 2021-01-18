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
	//Return uuid or error.
	Register(registration.User) (uuid.UUID, error)

	//Update user.
	//Return uuid or error.
	Update(registration.User) (uuid.UUID, error)

	//Get user by id.
	//Return user and error.
	Get(userId uuid.UUID) (registration.User, error)
}

type Resume interface {
	//Get resume by resume id.
	//Return resume and error.
	Get(resumeId uuid.UUID) (profile.Resume, error)

	//Get all resumes by user id.
	//Return slice of resumes and error.
	GetAll(userId uuid.UUID) ([]profile.Resume, error)

	//Create resume.
	//Return uuid or error.
	Create(resume profile.Resume) (uuid.UUID, error)

	//Update resume.
	//Return uuid or error.
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
	//Return slice of vacancy and error.
	GetWithPaging(userId uuid.UUID, count int, page int) ([]profile.Vacancy, error)

	//Create vacancy.
	//Return uuid or error.
	Create(vacancy profile.Vacancy) (uuid.UUID, error)

	//Update vacancy.
	//Return uuid or error.
	Update(vacancy profile.Vacancy) (uuid.UUID, error)

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
		User:    user.NewUserPostgres(db),
		Resume:  repository.NewResumePostgres(db),
		Vacancy: vacancy.NewVacancyPostgres(db),
	}
}
