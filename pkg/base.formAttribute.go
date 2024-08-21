package goFormValidators

import (
	"time"
)

type FieldAttribute struct {
	Name   string
	Code   string
	Errors []interface{}
}

type FieldAttributeInterface interface {
	// TODO: Add a method to setup default error message

	// AddError adds an error message to the attribute.
	AddError(message string)
	// GetCode returns the code of the attribute. Use to format the error message.
	GetCode() string
	// GetErrors returns the errors associated with the attribute.
	GetErrors() []interface{}
	// Time returns the time value of the attribute.
	Time() *time.Time
	// IsClean returns true if the attribute has no errors, false otherwise.
	IsClean() bool

	// Validators
	// ValidateRequired validates if the attribute is required.
	ValidateRequired()

	// ValidateMin validates if the attribute is greater than the min value.
	ValidateMin(min interface{}, message *string)
	// ValidateMax validates if the attribute is less than the max value.
	ValidateMax(max interface{}, message *string)
	// ValidateFormat validates if the attribute matches the format of regex. Using regexp.MustCompile & regexp.MatchString
	ValidateFormat(formatter string, formatterRemind string)
	// ValidateGt validates if the attribute is greater than the value.
	ValidateGt(value interface{}, message *string)
	// ValidateGtEq validates if the attribute is greater than or equal to the value.
	ValidateGtEq(value interface{}, message *string)
	// ValidateLt validates if the attribute is less than the value.
	ValidateLt(value interface{}, message *string)
	// ValidateLtEq validates if the attribute is less than or equal to the value.
	ValidateLtEq(value interface{}, message *string)
}

func (attribute *FieldAttribute) defaultValidateRequiredMessage() string {
	return "is required"
}
