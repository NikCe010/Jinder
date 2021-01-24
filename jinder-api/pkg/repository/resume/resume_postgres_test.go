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

	createItemCommand := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_type, "+
		"programmer_language) VALUES ($1,$2,$3,$4,$5)", "resumes")

	createWorkExpCommand := fmt.Sprintf("INSERT INTO %s (id, resume_id, company_name, experience_from, "+
		"experience_to, extra_content) VALUES ($1,$2,$3,$4,$5, $6)", "work_experience")

	mock.ExpectBegin()
	mock.ExpectExec(createItemCommand).
		WithArgs(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage).
		WillReturnResult(sqlmock.NewResult(0, 1))

	for _, w := range r.WorkExperiences {
		mock.ExpectExec(createWorkExpCommand).
			WithArgs(w.Id, w.ResumeId, w.CompanyName, w.From, w.To, w.Content).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}
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

	query := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language "+
		"FROM %s WHERE id = $1", "resumes")
	rows := sqlmock.NewRows([]string{"id", "user_id", "programmer_level", "programmer_type", "programmer_language"}).
		AddRow(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage)

	getAllQuery := fmt.Sprintf("SELECT id, resume_id, company_name, experience_from, experience_to, extra_content "+
		"FROM %s where resume_id = $1", "work_experience")

	mock.ExpectQuery(query).
		WithArgs(r.Id).
		WillReturnRows(rows)

	for _, w := range r.WorkExperiences {
		workExpRows := sqlmock.NewRows([]string{"id", "resume_id", "company_name", "experience_from", "experience_to", "extra_content"}).
			AddRow(w.Id, w.ResumeId, w.CompanyName, w.From, w.To, w.Content)
		mock.ExpectQuery(getAllQuery).
			WithArgs(r.Id).
			WillReturnRows(workExpRows)
	}

	user, err := repo.Get(r.Id)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestResumePostgres_GetWithPaging(t *testing.T) {
	db, mock := newMock()
	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	count := 10
	page := 0
	getAllQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language "+
		"FROM %s WHERE user_id = ?"+
		"OFFSET %d LIMIT %d", "resumes", page*count, count)

	rows := sqlmock.NewRows([]string{"id", "user_id", "programmer_level", "programmer_type", "programmer_language"}).
		AddRow(r.Id, r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage)

	mock.ExpectQuery(getAllQuery).
		WithArgs(r.UserId).
		WillReturnRows(rows)

	resume, err := repo.GetWithPaging(r.UserId, count, page)
	assert.NotNil(t, resume)
	assert.NoError(t, err)
}

func TestResumePostgres_Update(t *testing.T) {
	db, mock := newMock()

	repo := NewResumePostgres(db)
	defer func() {
		repo.db.Close()
	}()

	command := fmt.Sprintf("UPDATE %s SET user_id=$1, programmer_level=$2, programmer_type=$3, "+
		"programmer_language=$4 WHERE id=$5", "resumes")

	updateWorkExpCommand := fmt.Sprintf("UPDATE %s SET resume_id=$1, company_name=$2, experience_from=$3, "+
		"experience_to=$4, extra_content=$5 where id=$6", "work_experience")

	mock.ExpectBegin()
	mock.ExpectExec(command).
		WithArgs(r.UserId, r.ProgrammerLevel, r.ProgrammerType, r.ProgrammerLanguage, r.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	for _, w := range r.WorkExperiences {
		mock.ExpectExec(updateWorkExpCommand).
			WithArgs(w.ResumeId, w.CompanyName, w.From, w.To, w.Content, w.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}
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

	command := fmt.Sprintf("DELETE FROM %s WHERE id=$1", "resumes")
	deleteWorkExpCommand := fmt.Sprintf("DELETE FROM %s WHERE resume_id=$1", "work_experience")

	mock.ExpectBegin()
	mock.ExpectExec(command).
		WithArgs(r.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec(deleteWorkExpCommand).
		WithArgs(r.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.Delete(r.Id)

	assert.NoError(t, err)
}
