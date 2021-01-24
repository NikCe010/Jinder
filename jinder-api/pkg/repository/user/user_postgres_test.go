package repository

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var u = registration.User{
	Id: uuid.New(),
	Person: registration.Person{
		Name:     "Vasya",
		Surname:  "Vasyliev",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	Credentials: registration.Credentials{
		Email:        "test@mail.ru",
		PasswordHash: "$5$MnfsQ4iN$ZMTppKN16y/tIsUYs/obHlhdP.Os80yXhTurpBMUbA5",
	},
	Role: registration.Programmer,
}

func newMock() (*sqlx.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := sqlx.NewDb(sqldb, "sqlmock")
	return db, mock
}

func TestUserPostgres_Register(t *testing.T) {
	db, mock := newMock()
	repo := NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	registerQuery := fmt.Sprintf("INSERT INTO %s (id, name, surname, birthday, email, password_hash, role) VALUES ($1,$2,$3,$4,$5,$6,$7)", "users")

	mock.ExpectBegin()
	mock.ExpectPrepare(registerQuery).
		ExpectExec().
		WithArgs(u.Id, u.Name, u.Surname, u.Birthday, u.Email, u.PasswordHash, u.Role).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	user, err := repo.Register(u)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestUserPostgres_Get(t *testing.T) {
	db, mock := newMock()
	repo := NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	getItemQuery := fmt.Sprintf("SELECT id, name, surname, birthday, email, password_hash, role "+
		"FROM %s WHERE id = $1", "users")
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "birthday", "email", "password_hash", "role"}).
		AddRow(u.Id, u.Name, u.Surname, u.Birthday, u.Email, u.PasswordHash, u.Role)

	mock.ExpectQuery(getItemQuery).
		WithArgs(u.Id).
		WillReturnRows(rows)

	user, err := repo.Get(u.Id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserPostgres_Update(t *testing.T) {
	db, mock := newMock()
	repo := NewUserPostgres(db)

	db.Begin()

	defer func() {
		db.Close()
	}()

	updateQuery := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2, birthday=$3, email=$4, password_hash=$5, role=$6 WHERE id=$7", "users")

	mock.ExpectBegin()
	mock.ExpectPrepare(updateQuery).
		ExpectExec().
		WithArgs(u.Name, u.Surname, u.Birthday, u.Email, u.PasswordHash, u.Role, u.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	user, err := repo.Update(u)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
