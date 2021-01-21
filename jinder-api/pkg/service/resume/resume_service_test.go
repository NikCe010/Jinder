package resume

import (
	"Jinder/jinder-api/pkg/domain/profile"
	domain "Jinder/jinder-api/pkg/domain/profile/shared"
	"Jinder/jinder-api/pkg/service/dto"
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var (
	resumeId = uuid.New()
)

var resumeDto = dto.Resume{
	Id:                 resumeId,
	UserId:             uuid.New(),
	ProgrammerLevel:    shared.Middle,
	ProgrammerType:     shared.Backend,
	ProgrammerLanguage: shared.Golang,
	WorkExperiences: []dto.WorkExperience{
		{Id: uuid.New(),
			ResumeId:    resumeId,
			CompanyName: "TestCompany",
			From:        time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			To:          time.Now(),
			Content:     "Created a monitoring system, mentored two juniors, migrated the project to gRPS"},
	},
}

var resume = profile.Resume{
	Id:                 resumeId,
	UserId:             uuid.New(),
	ProgrammerLevel:    domain.Middle,
	ProgrammerType:     domain.Backend,
	ProgrammerLanguage: domain.Golang,
	WorkExperiences: []profile.WorkExperience{
		{Id: uuid.New(),
			ResumeId:    resumeId,
			CompanyName: "TestCompany",
			From:        time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			To:          time.Now(),
			Content:     "Created a monitoring system, mentored two juniors, migrated the project to gRPS"},
	},
}

type MockResumeRepository struct {
	mock.Mock
}

func (m MockResumeRepository) Get(resumeId uuid.UUID) (profile.Resume, error) {
	return resume, nil
}

func (m MockResumeRepository) GetAll(userId uuid.UUID) ([]profile.Resume, error) {
	return []profile.Resume{resume}, nil
}

func (m MockResumeRepository) Create(resume profile.Resume) (uuid.UUID, error) {
	return resume.Id, nil
}

func (m MockResumeRepository) Update(resume profile.Resume) (uuid.UUID, error) {
	return resume.Id, nil
}

func (m MockResumeRepository) Delete(resumeId uuid.UUID) error {
	return nil
}

func TestResumeService_Create(t *testing.T) {
	service := NewService(MockResumeRepository{})

	id, err := service.Create(resumeDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestResumeService_Get(t *testing.T) {
	service := NewService(MockResumeRepository{})

	result, err := service.Get(resumeDto.UserId)

	assert.NoError(t, err)
	assert.Equal(t, result.Id, resume.Id)
	assert.Equal(t, result.UserId, resume.UserId)
	assert.Equal(t, result.ProgrammerType, shared.ProgrammerType(resume.ProgrammerType))
	assert.Equal(t, result.ProgrammerLevel, shared.ProgrammerLevel(resume.ProgrammerLevel))
	assert.Equal(t, result.ProgrammerLanguage, shared.ProgrammerLanguage(resume.ProgrammerLanguage))
}

func TestResumeService_GetAll(t *testing.T) {
	service := NewService(MockResumeRepository{})

	resumes, err := service.GetAll(resumeDto.UserId)

	assert.NoError(t, err)
	assert.Equal(t, resumes[0].Id, resume.Id)
	assert.Equal(t, resumes[0].UserId, resume.UserId)
	assert.Equal(t, resumes[0].ProgrammerType, shared.ProgrammerType(resume.ProgrammerType))
	assert.Equal(t, resumes[0].ProgrammerLevel, shared.ProgrammerLevel(resume.ProgrammerLevel))
	assert.Equal(t, resumes[0].ProgrammerLanguage, shared.ProgrammerLanguage(resume.ProgrammerLanguage))
}

func TestResumeService_Update(t *testing.T) {
	service := NewService(MockResumeRepository{})

	id, err := service.Update(resumeDto)

	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.UUID{})
}

func TestResumeService_Delete(t *testing.T) {
	service := NewService(MockResumeRepository{})

	err := service.Delete(resumeDto.Id)

	assert.NoError(t, err)
}
