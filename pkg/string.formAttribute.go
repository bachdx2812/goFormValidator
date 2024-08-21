package goFormValidator

import (
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

// StringAttribute represents a string attribute validator.
type StringAttribute struct {
	FieldAttribute

	Value string
}

// GetCode returns the code of the string attribute.
func (attribute *StringAttribute) GetCode() string {
	return attribute.Code
}

// GetFieldCode returns the code of the bool attribute.
func (attribute *StringAttribute) GetFieldI18nCode() string {
	if attribute.Name != "" {
		return attribute.Name
	}
	return attribute.Code
}

// GetErrors returns the errors of the string attribute.
func (attribute *StringAttribute) GetErrors() []interface{} {
	return attribute.Errors
}

// AddError adds an error to the string attribute.
func (attribute *StringAttribute) AddError(message string) {
	attribute.Errors = append(attribute.Errors, message)
}

// ValidateRequired validates if the string attribute is required.
func (attribute *StringAttribute) ValidateRequired(message *string) {
	if attribute.Value == "" || strings.TrimSpace(attribute.Value) == "" {
		attribute.AddError(*message)
	}
}

func (attribute *StringAttribute) ValidateFormat(formatter string, formatterRemind string, message *string) {
	if attribute.Value != "" {
		re := regexp.MustCompile(formatter)

		if re.MatchString(attribute.Value) {
			return
		}

		attribute.AddError(*message)
	}
}

func (attribute *StringAttribute) Time() *time.Time {
	return nil
}

func (attribute *StringAttribute) IsClean() bool {
	return len(attribute.Errors) == 0
}

func (attribute *StringAttribute) ValidateMin(min interface{}, message *string) {
	switch v := min.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attribute.Value)) < v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *StringAttribute) ValidateMax(max interface{}, message *string) {
	switch v := max.(type) {
	case int64:
		if v < int64(utf8.RuneCountInString(attribute.Value)) {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *StringAttribute) ValidateGt(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attribute.Value)) <= v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *StringAttribute) ValidateGtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attribute.Value)) < v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *StringAttribute) ValidateLt(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attribute.Value)) >= v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}

func (attribute *StringAttribute) ValidateLtEq(value interface{}, message *string) {
	switch v := value.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attribute.Value)) > v {
			attribute.AddError(*message)
		}
	default:
		panic("Need to provide int64 interface{} as params")
	}
}
