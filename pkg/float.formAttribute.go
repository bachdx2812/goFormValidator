package goFormValidators

import (
	"server/app/pkg/translator"
	"time"

	"golang.org/x/exp/constraints"
)

type FloatAttribute[T constraints.Float] struct {
	FieldAttribute
	Value     T
	AllowZero bool
}

// GetCode returns the code of the attribute.
func (attribute *FloatAttribute[T]) GetCode() string {
	return attribute.Code
}

// GetFieldCode returns the code of the bool attribute.
func (attribute *FloatAttribute[T]) GetFieldI18nCode() string {
	if attribute.Name != "" {
		return attribute.Name
	}
	return attribute.Code
}

// GetErrors returns the errors associated with the attribute.
func (attribute *FloatAttribute[T]) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *FloatAttribute[T]) AddError(message string) {
	attribute.Errors = append(attribute.Errors, message)
}

// ValidateRequired validates if the attribute is required.
func (attribute *FloatAttribute[T]) ValidateRequired() {
	if attribute.Value == 0.0 && !attribute.AllowZero {
		attribute.AddError("is required")
	}
}

func (attribute *FloatAttribute[T]) ValidateFormat(formatter string, formatterRemind string) {
	// No need to implement yet
}

func (attribute *FloatAttribute[T]) Time() *time.Time {
	return nil
}

func (attribute *FloatAttribute[T]) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *FloatAttribute[T]) ValidateMin(min interface{}, message *string) {
	switch v := min.(type) {
	case float64:
		if float64(attribute.Value) < v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide float interface{} as params")
	}
}

func (attribute *FloatAttribute[T]) ValidateMax(max interface{}, message *string) {
	switch v := max.(type) {
	case float64:
		if v < float64(attribute.Value) {
			attribute.AddError(translator.Translate(nil, "ValidationFloatMax", max))
		}
	default:
		panic("Need to provide float64 interface{} as params")
	}
}

func (attribute *FloatAttribute[T]) ValidateGt(value interface{}, message *string) {
	switch v := value.(type) {
	case float64:
		if float64(attribute.Value) <= v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide float interface{} as params")
	}
}

func (attribute *FloatAttribute[T]) ValidateGtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case float64:
		if float64(attribute.Value) < v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide float interface{} as params")
	}
}

func (attribute *FloatAttribute[T]) ValidateLt(value interface{}, message *string) {
	switch v := value.(type) {
	case float64:
		if float64(attribute.Value) >= v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide float interface{} as params")
	}
}

func (attribute *FloatAttribute[T]) ValidateLtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case float64:
		if float64(attribute.Value) > v {

			attribute.AddError(*message)

		}
	default:
		panic("Need to provide float interface{} as params")
	}
}
