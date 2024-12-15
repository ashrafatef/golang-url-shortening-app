package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationError struct {
	Field      string `json:"field"`
	Constraint string `json:"constraint"`
}

func Validation(input interface{}) []ValidationError {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(input)
	var errors []ValidationError
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			logrus.Error(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			constraint := getConstraint(err)
			errors = append(errors, ValidationError{
				Field:      err.Field(),
				Constraint: constraint,
			})

		}
		return errors
	}
	return nil
}

func getConstraint(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "IsRequired"
	case "email":
		return "InvalidEmail"
	case "url":
		return "InvalidUrl"
	case "gte":
		return "Should be greater than " + fe.Param()
	case "lte":
		return "Should be lesser than " + fe.Param()
	default:
		return "Invalid field"
	}
}
