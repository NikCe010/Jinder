package mocks

import (
	"Jinder/jinder-api/jobs/pkg/infrastructure/event"
	"github.com/stretchr/testify/mock"
)

type MockRecommendationClient struct {
	mock.Mock
}

func (m MockRecommendationClient) NotifyWhenVacancyAdded(vacancyCreated event.VacancyCreated) error {
	return nil
}

func (m MockRecommendationClient) NotifyWhenResumeAdded(resumeCreated event.ResumeCreated) error {
	return nil
}

func (m MockRecommendationClient) NotifyWhenVacancyViewed(vacancyCreated event.VacancyViewed) error {
	return nil
}

func (m MockRecommendationClient) NotifyWhenResumeViewed(resumeCreated event.ResumeViewed) error {
	return nil
}
