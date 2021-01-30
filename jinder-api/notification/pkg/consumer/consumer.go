package consumer

import (
	"Jinder/jinder-api/recommendations/pkg/event"
	"Jinder/jinder-api/recommendations/pkg/service"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

type Consumer struct {
	Service *service.Service
}

func NewConsumer(service *service.Service) *Consumer {
	return &Consumer{Service: service}
}

func (c *Consumer) InitRoutes(url string) {
	nc, err := nats.Connect(url,
		nats.ErrorHandler(func(nc *nats.Conn, s *nats.Subscription, err error) {
			if s != nil {
				log.Printf("Async error in %q/%q: %v", s.Subject, s.Queue, err)
			} else {
				log.Printf("Async error outside subscription: %v", err)
			}
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)

	if _, err := ec.Subscribe("resume_created", func(s *event.ResumeCreated) {
		c.Service.ResumeCreatedEvent()
		log.Printf("New resume added")
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := ec.Subscribe("vacancy_created", func(s *event.VacancyCreated) {
		c.Service.VacancyCreatedEvent()
		log.Printf("New vacancy added")
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := ec.Subscribe("resume_viewed", func(s *event.ResumeViewed) {
		c.Service.ResumeViewedEvent()
		log.Printf("New resume added")
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := ec.Subscribe("vacancy_viewed", func(s *event.VacancyViewed) {
		c.Service.VacancyViewedEvent()
		log.Printf("New vacancy added")
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
