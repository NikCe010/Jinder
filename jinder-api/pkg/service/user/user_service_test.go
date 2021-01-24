package user

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"Jinder/jinder-api/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userDto = dto.User{
	Id: uuid.New(),
	Person: dto.Person{
		Name:     "Vasya",
		Surname:  "Vasyliev",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	Credentials: dto.Credentials{
		Email:                "testemail@mail.ru",
		Password:             "qwerty123",
		PasswordConfirmation: "qwerty123",
	},
	Role: dto.Programmer,
}

var userErrorDto = dto.User{
	Id: uuid.New(),
	Person: dto.Person{
		Name:     "Vasya",
		Surname:  "Vasyliev",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	Credentials: dto.Credentials{
		Email:                "testemail@mail.ru",
		Password:             "qwerty123",
		PasswordConfirmation: "qwe123rty",
	},
	Role: dto.Programmer,
}

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
	return registration.User{
		Id: userId,
		Person: registration.Person{
			Name:     "Vasya",
			Surname:  "Vasyliev",
			Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Credentials: registration.Credentials{
			Email:        "testemail@mail.ru",
			PasswordHash: "qwerty123",
		},
		Role: registration.Programmer,
	}, nil
}

func (m MockUserRepository) GetByEmail(email string) (registration.User, error) {
	return registration.User{
		Id: uuid.UUID{},
		Person: registration.Person{
			Name:     "Vasya",
			Surname:  "Vasyliev",
			Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Credentials: registration.Credentials{
			Email:        email,
			PasswordHash: "qwerty123",
		},
		Role: registration.Programmer,
	}, nil
}

func TestUserService_Register_WithValidData_ShouldCompleteSuccessful(t *testing.T) {
	service := NewService(MockUserRepository{})

	id, err := service.Register(userDto)

	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestUserService_Register_WithWrongPasswordConfirmation_ShouldReturnError(t *testing.T) {
	service := NewService(MockUserRepository{})

	id, err := service.Register(userErrorDto)

	assert.Error(t, err)
	assert.Equal(t, id, uuid.UUID{})
}

func TestUserService_Get(t *testing.T) {
	service := NewService(MockUserRepository{})

	user, err := service.GetUser(uuid.New())

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserService_Update_WithValidData_ShouldCompleteSuccessful(t *testing.T) {
	service := NewService(MockUserRepository{})

	user, err := service.UpdateUser(userDto)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserService_Update_WithWrongPasswordConfirmation_ShouldReturnError(t *testing.T) {
	service := NewService(MockUserRepository{})

	user, err := service.UpdateUser(userDto)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
