package repository

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"Jinder/jinder-api/pkg/domain/profile/shared"
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
	WorkExperiences: []profile.WorkExperience{
		{Id: uuid.New(),
			ResumeId:    resumeId,
			CompanyName: "TestCompany",
			From:        time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			To:          time.Now(),
			Content:     "Created a monitoring system, mentored two juniors, migrated the project to gRPS"},
	},
}

func newMock() (*sqlx.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := sqlx.NewDb(sqldb, "sqlmock")
	return db, mock
}

func TestResumePostgres_Create(t *testing.T) {
	db, mock := newMock()

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	query := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_type, "+
		"programmer_language) VALUES (?,?,?,?,?)", "resumes")

	mock.ExpectBegin()
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	id, err := repo.Create(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestResumePostgres_Get(t *testing.T) {
	db, mock := newMock()

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	query := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language FROM %s WHERE id = ?", "resumes")

	rows := sqlmock.NewRows([]string{"id", "user_id", "programmer_level", "programmer_type", "programmer_language"}).
		AddRow(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage)

	mock.ExpectQuery(query).
		WithArgs(r.Id).
		WillReturnRows(rows)

	user, err := repo.Get(r.Id)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestResumePostgres_GetAll(t *testing.T) {
	db, mock := newMock()

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	query := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language FROM %s WHERE user_id = ?", "resumes")

	rows := sqlmock.NewRows([]string{"id", "user_id", "programmer_level", "programmer_type", "programmer_language"}).
		AddRow(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage)

	mock.ExpectQuery(query).
		WithArgs(r.UserId).
		WillReturnRows(rows)

	resume, err := repo.GetAll(r.UserId)
	assert.NotNil(t, resume)
	assert.NoError(t, err)
}

func TestResumePostgres_Update(t *testing.T) {
	db, mock := newMock()

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	query := fmt.Sprintf("UPDATE %s SET user_id=?, programmer_level=?, programmer_type=?, "+
		"programmer_language=? WHERE id=?", "resumes")

	mock.ExpectBegin()
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs(r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage, r.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	id, err := repo.Update(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestResumePostgres_Delete(t *testing.T) {
	db, mock := newMock()

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", "resumes")

	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs(r.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(r.Id)

	assert.NoError(t, err)
}
