package models

import (
	"html/template"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Email struct {
	To      []string      `json:"to" validate:"required,dive,email"`
	From    string        `json:"from" validate:"required,email"`
	CC      []string      `json:"cc,omitempty" validate:"omitempty,dive,email"`
	BCC     []string      `json:"bcc,omitempty" validate:"omitempty,dive,email"`
	Subject string        `json:"subject" validate:"required,min=3,max=64"`
	Body    template.HTML `json:"body" validate:"required"`
}

func (r *Email) Bind(c *fiber.Ctx, v *validator.Validate) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Struct(r); err != nil {
		return err
	}

	return nil
}
