package gfv

import (
	"time"

	"golang.org/x/exp/constraints"
)

type IntAttribute[T constraints.Signed] struct {
	FieldAttribute
	Value     T
	AllowZero bool
}

// GetCode returns the code of the attribute.
func (attribute *IntAttribute[T]) GetCode() string {
	return attribute.Code
}

// GetFieldCode returns the code of the bool attribute.
func (attribute *IntAttribute[T]) GetFieldI18nCode() string {
	if attribute.Name != "" {
		return attribute.Name
	}
	return attribute.Code
}

// GetErrors returns the errors associated with the attribute.
func (attribute *IntAttribute[T]) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error message to the attribute.
func (attribute *IntAttribute[T]) AddError(message string) {
	attribute.Errors = append(attribute.Errors, message)
}

// ValidateRequired validates if the attribute is required.
func (attribute *IntAttribute[T]) ValidateRequired(message *string) {
	if attribute.Value == 0 && !attribute.AllowZero {
		attribute.AddError(*message)
	}
}

func (attribute *IntAttribute[T]) ValidateFormat(formatter string, formatterRemind string, message *string) {
	// No need to implement yet
}

func (attribute *IntAttribute[T]) Time() *time.Time {
	return nil
}

func (attribute *IntAttribute[T]) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *IntAttribute[T]) ValidateMin(min interface{}, message *string) {
	switch v := min.(type) {
	case int64:
		if int64(attribute.Value) < v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *IntAttribute[T]) ValidateMax(max interface{}, message *string) {
	switch v := max.(type) {
	case int64:
		if v < int64(attribute.Value) {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *IntAttribute[T]) ValidateGt(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(attribute.Value) <= v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *IntAttribute[T]) ValidateGtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(attribute.Value) < v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *IntAttribute[T]) ValidateLt(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(attribute.Value) >= v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *IntAttribute[T]) ValidateLtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(attribute.Value) > v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}
