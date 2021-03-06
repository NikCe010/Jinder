package repository

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	"Jinder/jinder-api/jobs/pkg/domain/registration"
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

	GetByEmail(email string) (registration.User, error)
}

type Resume interface {
	//Get resume by resume id.
	//Return resume and error.
	Get(resumeId uuid.UUID) (profile.Resume, error)

	//Get all resumes by user id.
	//Return slice of resumes and error.
	GetWithPaging(userId uuid.UUID, count, offset string) ([]profile.Resume, error)

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
	GetWithPaging(userId uuid.UUID, count, offset string) ([]profile.Vacancy, error)

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

type WorkExperience interface {
	//Get experience by vacancy id.
	//Return slice of experience and error.
	GetExperiences(vacancyId uuid.UUID) ([]profile.WorkExperience, error)

	//Create experience.
	//Return uuid or error.
	CreateExperience(experience profile.WorkExperience) (uuid.UUID, error)

	//Update experience.
	//Return uuid or error.
	UpdateExperience(experience profile.WorkExperience) (uuid.UUID, error)

	//Delete experience by vacancy id.
	//If failed return error.
	DeleteExperience(experienceId uuid.UUID) error
}

type Repository struct {
	User
	Resume
	Vacancy
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Resume:  NewResumePostgres(db),
		Vacancy: NewVacancyPostgres(db),
	}
}
