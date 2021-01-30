package infrastructure

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func NewNatsConnection(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer nc.Close()
	return nc, nil
}
