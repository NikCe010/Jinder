package tests

import (
	"Jinder/jinder-api/jobs/pkg/repository"
	"Jinder/jinder-api/jobs/pkg/tests/mocks"
	"Jinder/jinder-api/jobs/pkg/tests/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPostgres_Register(t *testing.T) {
	db, mock := mocks.NewSqlMock()
	repo := repository.NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	registerQuery := fmt.Sprintf("INSERT INTO %s (id, name, surname, birthday, email, password_hash, role) VALUES ($1,$2,$3,$4,$5,$6,$7)", "users")

	mock.ExpectBegin()
	mock.ExpectPrepare(registerQuery).
		ExpectExec().
		WithArgs(models.UserDomain.Id, models.UserDomain.Name, models.UserDomain.Surname, models.UserDomain.Birthday,
			models.UserDomain.Email, models.UserDomain.PasswordHash, models.UserDomain.Role).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	user, err := repo.Register(models.UserDomain)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestUserPostgres_Get(t *testing.T) {
	db, mock := mocks.NewSqlMock()
	repo := repository.NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	getItemQuery := fmt.Sprintf("SELECT id, name, surname, birthday, email, password_hash, role "+
		"FROM %s WHERE id = $1", "users")
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "birthday", "email", "password_hash", "role"}).
		AddRow(models.UserDomain.Id, models.UserDomain.Name, models.UserDomain.Surname, models.UserDomain.Birthday,
			models.UserDomain.Email, models.UserDomain.PasswordHash, models.UserDomain.Role)

	mock.ExpectQuery(getItemQuery).
		WithArgs(models.UserDomain.Id).
		WillReturnRows(rows)

	user, err := repo.Get(models.UserDomain.Id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserPostgres_GetByEmail(t *testing.T) {
	db, mock := mocks.NewSqlMock()
	repo := repository.NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	getQuery := fmt.Sprintf("SELECT id, name, surname, birthday, email, password_hash, role "+
		"FROM %s WHERE email = $1", "users")
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "birthday", "email", "password_hash", "role"}).
		AddRow(models.UserDomain.Id, models.UserDomain.Name, models.UserDomain.Surname, models.UserDomain.Birthday,
			models.UserDomain.Email, models.UserDomain.PasswordHash, models.UserDomain.Role)

	mock.ExpectQuery(getQuery).
		WithArgs(models.UserDomain.Email).
		WillReturnRows(rows)

	user, err := repo.GetByEmail(models.UserDomain.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserPostgres_Update(t *testing.T) {
	db, mock := mocks.NewSqlMock()
	repo := repository.NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	updateQuery := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2, birthday=$3, email=$4, password_hash=$5, role=$6 WHERE id=$7", "users")

	mock.ExpectBegin()
	mock.ExpectPrepare(updateQuery).
		ExpectExec().
		WithArgs(models.UserDomain.Name, models.UserDomain.Surname, models.UserDomain.Birthday, models.UserDomain.Email,
			models.UserDomain.PasswordHash, models.UserDomain.Role, models.UserDomain.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	user, err := repo.Update(models.UserDomain)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
