package engine

import (
	"reflect"
)

// everything is possible with reflect
func CopyInterface[T any](src T) T {
	origVal := reflect.ValueOf(src).Elem()

	newVal := reflect.New(origVal.Type()).Elem()

	newVal.Set(origVal)

	return newVal.Addr().Interface().(T)
}
