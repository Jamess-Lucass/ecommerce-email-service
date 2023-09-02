package handlers

import (
	"github.com/Jamess-Lucass/ecommerce-email-service/responses"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Start() error {
	f := fiber.New()

	f.Use(fiberzap.New(fiberzap.Config{
		Logger: s.logger,
	}))

	f.Get("/api/healthz", s.Healthz)

	f.Post("/api/v1/emails", s.SendEmail)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Code: fiber.StatusNotFound, Message: "No resource found"})
	})

	return f.Listen(":8080")
}
