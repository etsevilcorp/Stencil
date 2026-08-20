package stencil

import (
	"fmt"
	"reflect"
)

// use errors.AsType to check error(gatekeeping it for 1.26+)
type ValidationError struct {
	Field string
	Given string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("field %v contains unknown value %q", e.Field, e.Given)
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
			return &ValidationError{
				Field: t.Field(i).Tag.Get("toml"),
				Given: fmt.Sprintf("%v", enum),
			}
		}
	}

	return nil
}
