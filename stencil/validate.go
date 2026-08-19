package stencil

import "fmt"

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
	if !s.Anchor.Valid() {
		return &ValidationError{
			Field: "anchor",
			Given: string(s.Anchor),
		}
	}
	if !s.MaxHandling.Valid() {
		return &ValidationError{
			Field: "max-handling",
			Given: string(s.MaxHandling),
		}
	}
	if !s.MinHandling.Valid() {
		return &ValidationError{
			Field: "min-handling",
			Given: string(s.MinHandling),
		}
	}

	return nil
}
