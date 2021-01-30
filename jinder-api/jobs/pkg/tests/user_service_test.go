package tests

import (
	"Jinder/jinder-api/jobs/pkg/service"
	"Jinder/jinder-api/jobs/pkg/tests/mocks"
	"Jinder/jinder-api/jobs/pkg/tests/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Register_WithValidData_ShouldCompleteSuccessful(t *testing.T) {
	service := service.NewUserService(mocks.MockUserRepository{})

	id, err := service.Register(models.UserDto)

	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestUserService_Register_WithWrongPasswordConfirmation_ShouldReturnError(t *testing.T) {
	service := service.NewUserService(mocks.MockUserRepository{})

	id, err := service.Register(models.UserErrorDto)

	assert.Error(t, err)
	assert.Equal(t, id, uuid.UUID{})
}

func TestUserService_Get(t *testing.T) {
	service := service.NewUserService(mocks.MockUserRepository{})

	user, err := service.GetUser(uuid.New())

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserService_Update_WithValidData_ShouldCompleteSuccessful(t *testing.T) {
	service := service.NewUserService(mocks.MockUserRepository{})

	user, err := service.UpdateUser(models.UserDto)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserService_Update_WithWrongPasswordConfirmation_ShouldReturnError(t *testing.T) {
	service := service.NewUserService(mocks.MockUserRepository{})

	user, err := service.UpdateUser(models.UserDto)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
