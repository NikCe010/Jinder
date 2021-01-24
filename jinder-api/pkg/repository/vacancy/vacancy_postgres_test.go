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
)

var v = profile.Vacancy{
	Id:                 uuid.New(),
	UserId:             uuid.New(),
	ProgrammerLevel:    shared.Middle,
	ProgrammerLanguage: shared.Golang,
	ProgrammerType:     shared.Backend,
	CompanyName:        "Test Best Company",
	SalaryFrom:         "150000",
	SalaryTo:           "200000",
	ExtraBenefits:      "Medical Insurance, paid vacation 31 days",
}

func newMock() (*sqlx.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := sqlx.NewDb(sqldb, "sqlmock")
	return db, mock
}

func TestVacancyPostgres_Create(t *testing.T) {
	db, mock := newMock()
	repo := NewVacancyPostgres(db)

	defer func() {
		db.Close()
	}()

	createItemQuery := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_language, "+
		"programmer_type, company_name, salary_from, salary_to, extra_benefits) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)", "vacancies")

	mock.ExpectBegin()
	mock.ExpectPrepare(createItemQuery).
		ExpectExec().
		WithArgs(v.Id, v.UserId, v.ProgrammerLevel, v.ProgrammerLanguage, v.ProgrammerType, v.CompanyName,
			v.SalaryFrom, v.SalaryTo, v.ExtraBenefits).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	result, err := repo.Create(v)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestVacancyPostgres_Update(t *testing.T) {
	db, mock := newMock()
	repo := NewVacancyPostgres(db)

	defer func() {
		db.Close()
	}()

	updateItemQuery := fmt.Sprintf("UPDATE %s SET user_id=$1, programmer_level=$2, programmer_language=$3, "+
		"programmer_type=$4, company_name=$5, salary_from=$6, salary_to=$7, extra_benefits=$8 WHERE id=$9", "vacancies")

	mock.ExpectBegin()
	mock.ExpectPrepare(updateItemQuery).
		ExpectExec().
		WithArgs(v.UserId, v.ProgrammerLevel, v.ProgrammerLanguage, v.ProgrammerType, v.CompanyName,
			v.SalaryFrom, v.SalaryTo, v.ExtraBenefits, v.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	result, err := repo.Update(v)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestVacancyPostgres_Get(t *testing.T) {
	db, mock := newMock()
	repo := NewVacancyPostgres(db)

	defer func() {
		db.Close()
	}()

	getItemQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_language, "+
		"programmer_type, company_name, salary_from, salary_to, extra_benefits FROM %s WHERE id=$1", "vacancies")

	rows := sqlmock.NewRows([]string{"id", "user_id", "programmer_level", "programmer_language", "programmer_type", "company_name", "salary_from", "salary_to", "extra_benefits"}).
		AddRow(v.Id, v.UserId, v.ProgrammerLevel, v.ProgrammerLanguage, v.ProgrammerType, v.CompanyName, v.SalaryFrom, v.SalaryTo, v.ExtraBenefits)

	mock.ExpectQuery(getItemQuery).
		WithArgs(v.Id).
		WillReturnRows(rows)

	result, err := repo.Get(v.Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestVacancyPostgres_GetWithPaging(t *testing.T) {
	db, mock := newMock()
	repo := NewVacancyPostgres(db)

	defer func() {
		db.Close()
	}()

	count := 10
	page := 0
	getItemsQuery := fmt.Sprintf("SELECT id, user_id, programmer_type, programmer_level, "+
		"programmer_language, company_name, salary_from, salary_to, extra_benefits FROM %s WHERE user_id=$1 "+
		"OFFSET %d LIMIT %d", "vacancies", page*count, count)
	rows := sqlmock.NewRows([]string{"id", "user_id", "programmer_type", "programmer_level", "programmer_language", "company_name", "salary_from", "salary_to", "extra_benefits"}).
		AddRow(v.Id, v.UserId, v.ProgrammerLevel, v.ProgrammerLanguage, v.ProgrammerType, v.CompanyName, v.SalaryFrom, v.SalaryTo, v.ExtraBenefits)

	mock.ExpectQuery(getItemsQuery).
		WithArgs(v.UserId).
		WillReturnRows(rows)

	result, err := repo.GetWithPaging(v.UserId, count, page)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestVacancyPostgres_Delete(t *testing.T) {
	db, mock := newMock()
	repo := NewVacancyPostgres(db)

	defer func() {
		db.Close()
	}()

	deleteItemQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", "vacancies")

	mock.ExpectPrepare(deleteItemQuery).
		ExpectExec().
		WithArgs(v.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(v.Id)

	assert.NoError(t, err)
}
