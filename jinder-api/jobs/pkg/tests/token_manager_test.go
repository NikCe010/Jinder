package tests

import (
	"Jinder/jinder-api/jobs/pkg/service"
	"Jinder/jinder-api/jobs/pkg/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenManagerService_Generate_WithValidPasswordAndEmail_ShouldSucceed(t *testing.T) {
	service := service.NewTokenManagerService(mocks.MockUserRepository{})

	token, err := service.Generate("test@test.ru", "123qwe123")

	assert.NoError(t, err)
	assert.NotEqual(t, token, "")
}

func TestTokenManagerService_Generate_WithWrongEmail_ShouldFailed(t *testing.T) {
	service := service.NewTokenManagerService(mocks.MockUserRepository{})

	token, err := service.Generate("tesqwet@test.ru", "123qwe123")

	assert.Error(t, err)
	assert.Equal(t, token, "")
}

func TestTokenManagerService_Generate_WithWrongPassword_ShouldFailed(t *testing.T) {
	service := service.NewTokenManagerService(mocks.MockUserRepository{})

	token, err := service.Generate("test@test.ru", "234qwe567")

	assert.Error(t, err)
	assert.Equal(t, token, "")
}

func TestTokenManagerService_Validate(t *testing.T) {
	service := service.NewTokenManagerService(mocks.MockUserRepository{})

	token, _ := service.Generate("test@test.ru", "123qwe123")
	userId, err := service.Validate(token)

	assert.NoError(t, err)
	assert.NotEqual(t, userId, uuid.UUID{})
}
