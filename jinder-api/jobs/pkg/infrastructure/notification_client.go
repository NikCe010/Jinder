package infrastructure

import (
	"Jinder/jinder-api/jobs/pkg/infrastructure/event"
	nats "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

type RecommendationClient struct {
	Nats *nats.Conn
}

func (e *RecommendationClient) NotifyWhenVacancyAdded(vacancyCreated event.VacancyCreated) error {
	ec, err := nats.NewEncodedConn(e.Nats, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
		return err
	}
	defer ec.Close()

	if err := ec.Publish("vacancy_created", &vacancyCreated); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *RecommendationClient) NotifyWhenResumeAdded(resumeCreated event.ResumeCreated) error {
	ec, err := nats.NewEncodedConn(e.Nats, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
		return err
	}
	defer ec.Close()

	if err := ec.Publish("resume_created", &resumeCreated); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *RecommendationClient) NotifyWhenVacancyViewed(vacancyViewed event.VacancyViewed) error {
	ec, err := nats.NewEncodedConn(e.Nats, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
		return err
	}
	defer ec.Close()

	if err := ec.Publish("vacancy_viewed", &vacancyViewed); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *RecommendationClient) NotifyWhenResumeViewed(resumeViewed event.ResumeViewed) error {
	ec, err := nats.NewEncodedConn(e.Nats, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
		return err
	}
	defer ec.Close()

	if err := ec.Publish("resume_viewed", &resumeViewed); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func NewRecommendationClient(conn *nats.Conn) *RecommendationClient {
	return &RecommendationClient{Nats: conn}
}
