package services

import (
	"fmt"
	"net/url"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

var user = os.Getenv("RABBITMQ_USERNAME")
var pass = os.Getenv("RABBITMQ_PASSWORD")
var host = os.Getenv("RABBITMQ_HOST")
var port = os.Getenv("RABBITMQ_PORT")

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Ping() error {
	u := &url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(user, pass),
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	rabbitMQClient, err := amqp091.Dial(u.String())
	if err != nil {
		return err
	}

	defer rabbitMQClient.Close()

	return nil
}
