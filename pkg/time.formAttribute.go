package gfv

import (
	"fmt"
	"server/app/pkg/translator"
	"time"
)

type TimeAttribute struct {
	FieldAttribute
	Value     string
	TimeValue *time.Time
}

// GetCode returns the code of the attribute.
func (attribute *TimeAttribute) GetCode() string {
	return attribute.Code
}

// GetErrors returns the errors of the string attribute.
func (attribute *TimeAttribute) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error to the string attribute.
func (attribute *TimeAttribute) AddError(message string) {
	attribute.Errors = append(attribute.Errors, message)
}

// ValidateRequired validates if the string attribute is required.
func (attribute *TimeAttribute) ValidateRequired(message *string) {
	if attribute.Value == "" {
		attribute.AddError(*message)
	}
}

func (attribute *TimeAttribute) ValidateFormat(formatter string, formatterRemind string, message *string) {
	if attribute.Value != "" {
		if timeValue, err := time.ParseInLocation(formatter, attribute.Value, time.Local); err != nil {
			attribute.AddError(*message)
		} else {
			attribute.TimeValue = &timeValue
		}
	}
}

func (attribute *TimeAttribute) Time() *time.Time {
	return attribute.TimeValue
}

func (attribute *TimeAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *TimeAttribute) ValidateMin(min interface{}, message *string) {
	switch v := min.(type) {
	case time.Time:
		if attribute.TimeValue != nil && v.After(*attribute.TimeValue) {
			attribute.AddError(translator.Translate(nil, "ValidationTimeMin", fmt.Sprintf("%+v", v)))
		}
	default:
		panic("Need to provide time interface{} as params")
	}
}

func (attribute *TimeAttribute) ValidateMax(max interface{}, message *string) {
	switch v := max.(type) {
	case time.Time:
		if attribute.TimeValue != nil && v.Before(*attribute.TimeValue) {
			attribute.AddError(translator.Translate(nil, "ValidationTimeMax", fmt.Sprintf("%+v", v)))
		}
	default:
		panic("Need to provide time interface{} as params")
	}
}

func (attribute *TimeAttribute) ValidateGt(value interface{}, message *string) {
	switch v := value.(type) {
	case time.Time:
		if attribute.TimeValue == nil || attribute.TimeValue.After(v) {
			return
		}

		attribute.AddError(*message)
	default:
		panic("Need to provide time interface{} as params")
	}
}

func (attribute *TimeAttribute) ValidateGtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case time.Time:
		if attribute.TimeValue != nil && v.After(*attribute.TimeValue) {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide time interface{} as params")
	}
}

func (attribute *TimeAttribute) ValidateLt(value interface{}, message *string) {
	switch v := value.(type) {
	case time.Time:
		if attribute.TimeValue == nil || attribute.TimeValue.Before(v) {
			return
		}

		attribute.AddError(*message)
	default:
		panic("Need to provide time interface{} as params")
	}
}

func (attribute *TimeAttribute) ValidateLtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case time.Time:
		if attribute.TimeValue != nil && v.Before(*attribute.TimeValue) {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide time interface{} as params")
	}
}
