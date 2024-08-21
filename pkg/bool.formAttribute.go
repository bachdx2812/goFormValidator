package goFormValidators

import (
	"time"
)

// BoolAttribute represents a bool attribute validator.
type BoolAttribute struct {
	FieldAttribute

	Value bool
}

// GetCode returns the code of the bool attribute.
func (attribute *BoolAttribute) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors of the bool attribute.
func (attribute *BoolAttribute) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error to the bool attribute.
func (attribute *BoolAttribute) AddError(message string) {
	attribute.Errors = append(attribute.Errors, message)
}

// ValidateRequired validates if the bool attribute is required.
func (attribute *BoolAttribute) ValidateRequired() {
}

func (attribute *BoolAttribute) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) Time() *time.Time {
	return nil
}

func (attribute *BoolAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *BoolAttribute) ValidateMin(min interface{}, message *string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) ValidateMax(max interface{}, message *string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) ValidateGt(value interface{}, message *string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) ValidateGtEq(value interface{}, message *string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) ValidateLt(value interface{}, message *string) {
	// No need to implement yet
}

func (attribute *BoolAttribute) ValidateLtEq(value interface{}, message *string) {
	// No need to implement yet
}
