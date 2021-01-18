package repository

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"Jinder/jinder-api/pkg/domain/profile/shared"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var (
	resumeId = uuid.New()
)

var r = profile.Resume{
	Id:                 resumeId,
	UserId:             uuid.New(),
	ProgrammerLevel:    shared.Middle,
	ProgrammerType:     shared.Backend,
	ProgrammerLanguage: shared.Golang,
	ExtraSkills:        []string{"RabbitMQ", "Redis", "Docker", "PostgreSQL", "MongoDB"},
	WorkExperiences: []profile.WorkExperience{
		{Id: uuid.New(),
			ResumeId:    resumeId,
			CompanyName: "TestCompany",
			From:        time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			To:          time.Now(),
			Content:     "Created a monitoring system, mentored two juniors, migrated the project to gRPS"},
	},
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestPostgres_Create(t *testing.T) {
	sqldb, mock := NewMock()
	db := sqlx.NewDb(sqldb, "sqlmock")

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	query := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_type, "+
		"programmer_language, extra_skills) VALUES (?,?,?,?,?,?)", "resumes")

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage, r.ExtraSkillsToText()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	id, err := repo.Create(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}
