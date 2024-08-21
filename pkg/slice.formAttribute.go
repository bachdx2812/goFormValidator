package goFormValidators

import (
	"server/app/gql/inputs/adminInputs"
	"server/app/gql/inputs/employerInputs"
	"server/app/gql/inputs/workerInputs"
	"time"
)

type NestedSlices interface {
	workerInputs.WorkerSkillInput |
		workerInputs.WorkerCertificationInput |
		employerInputs.JobCertificationInput |
		adminInputs.JobTypeUpdateSkillLevelInput |
		string
}

type SliceAttribute[T NestedSlices] struct {
	FieldAttribute
	Value *[]T
}

// GetCode returns the code of the attribute.
func (attribute *SliceAttribute[T]) GetCode() string {
	return attribute.Code
}

// GetFieldCode returns the code of the bool attribute.
func (attribute *SliceAttribute[T]) GetFieldI18nCode() string {
	if attribute.Name != "" {
		return attribute.Name
	}
	return attribute.Code
}

// GetErrors returns the errors associated with the attribute.
func (attribute *SliceAttribute[T]) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *SliceAttribute[T]) AddError(message string) {
	attribute.Errors = append(attribute.Errors, message)
}

// ValidateRequired validates if the attribute is required.
func (attribute *SliceAttribute[T]) ValidateRequired(message *string) {
	if attribute.Value == nil || len(*attribute.Value) == 0 {
		attribute.AddError(*message)
	}
}

func (attribute *SliceAttribute[T]) ValidateFormat(formatter string, formatterRemind string, message *string) {
	// No need to implement
}

func (attribute *SliceAttribute[T]) Time() *time.Time {
	return nil
}

func (attribute *SliceAttribute[T]) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *SliceAttribute[T]) ValidateMin(min interface{}, message *string) {
}

func (attribute *SliceAttribute[T]) ValidateMax(min interface{}, message *string) {
}

func (attribute *SliceAttribute[T]) ValidateGt(value interface{}, message *string) {}

func (attribute *SliceAttribute[T]) ValidateGtEq(value interface{}, message *string) {}

func (attribute *SliceAttribute[T]) ValidateLt(value interface{}, message *string) {}

func (attribute *SliceAttribute[T]) ValidateLtEq(value interface{}, message *string) {}
