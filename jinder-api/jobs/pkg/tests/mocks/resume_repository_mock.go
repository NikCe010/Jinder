package mocks

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	"Jinder/jinder-api/jobs/pkg/tests/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockResumeRepository struct {
	mock.Mock
}

func (m MockResumeRepository) Get(resumeId uuid.UUID) (profile.Resume, error) {
	return models.Resume, nil
}

func (m MockResumeRepository) GetWithPaging(userId uuid.UUID, count string, page string) ([]profile.Resume, error) {
	return []profile.Resume{models.Resume}, nil
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
