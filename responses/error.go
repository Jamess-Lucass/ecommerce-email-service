package responses

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code    uint    `json:"code"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

type Error struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

func NewErrorResponse(err error) interface{} {
	e := ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errors []Error
		for _, err := range validationErrors {
			element := Error{
				Field:    err.Field(),
				Messages: []string{getErrorMessageForValidationTag(err)},
			}
			errors = append(errors, element)
		}
		e.Message = "Failed to validate body"
		e.Errors = errors
	}

	return e
}

func getErrorMessageForValidationTag(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "min":
		if err.Kind() == reflect.Slice {
			return fmt.Sprintf("This field must contain greater than %s element(s)", err.Param())
		}
		return fmt.Sprintf("This field must be greater than %s character(s)", err.Param())
	case "max":
		if err.Kind() == reflect.Slice {
			return fmt.Sprintf("This field must contain less than %s element(s)", err.Param())
		}
		return fmt.Sprintf("This field must be less than than %s character(s)", err.Param())
	default:
		return err.Param()
	}
}
