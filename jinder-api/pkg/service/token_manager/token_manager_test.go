package token_manager

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockUserRepository struct {
	mock.Mock
}

func (m MockUserRepository) Register(user registration.User) (uuid.UUID, error) {
	return user.Id, nil
}

func (m MockUserRepository) Update(user registration.User) (uuid.UUID, error) {
	return user.Id, nil
}

func (m MockUserRepository) Get(userId uuid.UUID) (registration.User, error) {
	return registration.User{}, nil
}

func (m MockUserRepository) GetByEmail(email string) (registration.User, error) {
	if email == "test@test.ru" {
		id, _ := uuid.Parse("25790180-5c2d-11eb-ae93-0242ac131111")
		return registration.User{
			Id: id,
			Person: registration.Person{
				Name:     "Vasya",
				Surname:  "Vasyliev",
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Credentials: registration.Credentials{
				Email:        email,
				PasswordHash: "$2a$04$w8/XJrPLfa9K6uDdxynzc.eCL/kA2Yuonu8Y.1lZAZNr0N.OXdGTO",
			},
			Role: registration.Programmer,
		}, nil
	} else {
		return registration.User{
			Id: uuid.UUID{},
			Person: registration.Person{
				Name:     "Vasya",
				Surname:  "Vasyliev",
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Credentials: registration.Credentials{
				Email:        email,
				PasswordHash: "$2a$04$v89aJOj.BDmu8e/ict5IOXDHrigEZHsFq0/yjw5gXk6KUe7vVBP.",
			},
			Role: registration.Programmer,
		}, nil
	}

}

func TestTokenManagerService_Generate_WithValidPasswordAndEmail_ShouldSucceed(t *testing.T) {
	service := NewTokenManagerService(MockUserRepository{})

	token, err := service.Generate("test@test.ru", "123qwe123")

	assert.NoError(t, err)
	assert.NotEqual(t, token, "")
}

func TestTokenManagerService_Generate_WithWrongEmail_ShouldFailed(t *testing.T) {
	service := NewTokenManagerService(MockUserRepository{})

	token, err := service.Generate("tesqwet@test.ru", "123qwe123")

	assert.Error(t, err)
	assert.Equal(t, token, "")
}

func TestTokenManagerService_Generate_WithWrongPassword_ShouldFailed(t *testing.T) {
	service := NewTokenManagerService(MockUserRepository{})

	token, err := service.Generate("test@test.ru", "234qwe567")

	assert.Error(t, err)
	assert.Equal(t, token, "")
}

func TestTokenManagerService_Validate(t *testing.T) {
	service := NewTokenManagerService(MockUserRepository{})

	userId, err := service.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTE2MDA3MTIsImlhdCI6MTYxMTUxNDMxMiwidXNlcl9pZCI6IjI1NzkwMTgwLTVjMmQtMTFlYi1hZTkzLTAyNDJhYzEzMTExMSJ9.ofh4j72hc-Hz8il4rFW36jNuyv5sjI2c1M82ZQ3HkeE")

	assert.NoError(t, err)
	assert.NotEqual(t, userId, uuid.UUID{})
}
