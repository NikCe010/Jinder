package service

import "Jinder/jinder-api/recommendations/pkg/infrastructure"

type NotificationService struct {
	jobClient infrastructure.Jobs
}

func (n NotificationService) ResumeCreatedEvent() {
	panic("implement me")
}

func (n NotificationService) ResumeViewedEvent() {
	panic("implement me")
}

func (n NotificationService) VacancyCreatedEvent() {
	panic("implement me")
}

func (n NotificationService) VacancyViewedEvent() {
	panic("implement me")
}

func NewNotificationService(jobClient infrastructure.Jobs) *NotificationService {
	return &NotificationService{jobClient: jobClient}
}
