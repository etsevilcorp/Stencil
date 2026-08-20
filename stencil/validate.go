package stencil

import (
	"fmt"
	"reflect"
	"strconv"
)

// use errors.AsType to check error(gatekeeping it for 1.26+)
type UnknownError struct {
	Field string
	Given string
}

func (e *UnknownError) Error() string {
	return fmt.Sprintf("field %v contains unknown value %q", e.Field, e.Given)
}

type BadValueError struct {
	Field   string
	Given   string
	Message string
}

func (e *BadValueError) Error() string {
	return fmt.Sprintf("field %v contains value %q which didn't validate correctly; %v", e.Field, e.Given, e.Message)
}

// validates all fields. If validation didn't succeed returns an error
func (s Stencil) Validate() error {
	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		enum, ok := v.Field(i).Interface().(Enum)
		if !ok {
			continue
		}
		if !enum.Valid() {
			return &UnknownError{
				Field: t.Field(i).Tag.Get("toml"),
				Given: fmt.Sprintf("%v", enum),
			}
		}
	}

	if (s.MaxHeight != nil && s.MinHeight != nil) &&
		(*s.MaxHeight < *s.MinHeight) {
		return &BadValueError{
			Field:   "min/max-height",
			Given:   fmt.Sprintf("%v, %v", strconv.Itoa(*s.MinHeight), strconv.Itoa(*s.MaxHeight)),
			Message: "minimal value can't be bigger than maximum",
		}
	}

	if (s.MaxWidth != nil && s.MinWidth != nil) &&
		(*s.MaxWidth < *s.MinWidth) {
		return &BadValueError{
			Field:   "min/max-width",
			Given:   fmt.Sprintf("%v, %v", strconv.Itoa(*s.MinWidth), strconv.Itoa(*s.MaxWidth)),
			Message: "minimal value can't be bigger than maximum",
		}
	}

	return nil
}
