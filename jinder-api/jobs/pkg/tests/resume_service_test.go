package tests

import (
	"Jinder/jinder-api/jobs/pkg/service"
	"Jinder/jinder-api/jobs/pkg/service/dto/shared"
	"Jinder/jinder-api/jobs/pkg/tests/mocks"
	"Jinder/jinder-api/jobs/pkg/tests/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResumeService_Create(t *testing.T) {
	service := service.NewResumeService(mocks.MockResumeRepository{}, mocks.MockRecommendationClient{})

	id, err := service.CreateResume(models.ResumeDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestResumeService_Get(t *testing.T) {
	service := service.NewResumeService(mocks.MockResumeRepository{}, mocks.MockRecommendationClient{})

	result, err := service.GetResume(models.ResumeDto.UserId)

	assert.NoError(t, err)
	assert.Equal(t, result.Id, models.Resume.Id)
	assert.Equal(t, result.UserId, models.Resume.UserId)
	assert.Equal(t, result.ProgrammerType, shared.ProgrammerType(models.Resume.ProgrammerType))
	assert.Equal(t, result.ProgrammerLevel, shared.ProgrammerLevel(models.Resume.ProgrammerLevel))
	assert.Equal(t, result.ProgrammerLanguage, shared.ProgrammerLanguage(models.Resume.ProgrammerLanguage))
}

func TestResumeService_GetAll(t *testing.T) {
	service := service.NewResumeService(mocks.MockResumeRepository{}, mocks.MockRecommendationClient{})

	resumes, err := service.GetResumes(models.ResumeDto.UserId, "10", "0")

	assert.NoError(t, err)
	assert.Equal(t, resumes[0].Id, models.Resume.Id)
	assert.Equal(t, resumes[0].UserId, models.Resume.UserId)
	assert.Equal(t, resumes[0].ProgrammerType, shared.ProgrammerType(models.Resume.ProgrammerType))
	assert.Equal(t, resumes[0].ProgrammerLevel, shared.ProgrammerLevel(models.Resume.ProgrammerLevel))
	assert.Equal(t, resumes[0].ProgrammerLanguage, shared.ProgrammerLanguage(models.Resume.ProgrammerLanguage))
}

func TestResumeService_Update(t *testing.T) {
	service := service.NewResumeService(mocks.MockResumeRepository{}, mocks.MockRecommendationClient{})

	id, err := service.UpdateResume(models.ResumeDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestResumeService_Delete(t *testing.T) {
	service := service.NewResumeService(mocks.MockResumeRepository{}, mocks.MockRecommendationClient{})

	err := service.DeleteResume(models.ResumeDto.Id)

	assert.NoError(t, err)
}
