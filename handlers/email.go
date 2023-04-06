package handlers

import (
	"github.com/Jamess-Lucass/ecommerce-email-service/models"
	"github.com/Jamess-Lucass/ecommerce-email-service/responses"
	"github.com/gofiber/fiber/v2"
)

// POST: /api/v1/emails
func (s *Server) SendEmail(c *fiber.Ctx) error {

	req := &models.Email{}
	if err := req.Bind(c, s.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err))
	}

	if err := s.emailService.AddToQueue(c.Context(), *req); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(req)
}
