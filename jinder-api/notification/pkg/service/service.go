package service

import (
	"Jinder/jinder-api/recommendations/pkg/infrastructure"
)

type Notification interface {
	ResumeCreatedEvent()
	ResumeViewedEvent()
	VacancyCreatedEvent()
	VacancyViewedEvent()
}

type Service struct {
	Notification
}

func NewService(infrastructure *infrastructure.Infrastructure) *Service {
	return &Service{
		Notification: NewNotificationService(infrastructure),
	}
}
