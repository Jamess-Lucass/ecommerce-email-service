package services

import (
	"encoding/json"

	"github.com/Jamess-Lucass/ecommerce-email-service/models"
	"github.com/rabbitmq/amqp091-go"
	"github.com/valyala/fasthttp"
)

type EmailService struct {
	rabbitMQ *amqp091.Channel
}

func NewEmailService(rabbitMQ *amqp091.Channel) *EmailService {
	return &EmailService{rabbitMQ: rabbitMQ}
}

func (s *EmailService) AddToQueue(ctx *fasthttp.RequestCtx, email models.Email) error {
	q, err := s.rabbitMQ.QueueDeclare(
		"emails", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(email)
	if err != nil {
		return err
	}

	if err := s.rabbitMQ.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			Body: body,
		}); err != nil {
		return err
	}

	return nil
}
