package assert

import (
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v; want %v", got, want)
	}
}

// shamelessly lifted from stretchr/testify
func isNil(got any) bool {
	if got == nil {
		return true
	}
	value := reflect.ValueOf(got)
	switch value.Kind() {
	case
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:

		return value.IsNil()
	default:
		return false
	}
}

func IsNil(t *testing.T, got any) {
	t.Helper()
	if !isNil(got) {
		t.Errorf("'%v' is not nil", got)
	}
}

func IsNotNil(t *testing.T, got any) {
	t.Helper()
	if isNil(got) {
		t.Errorf("'%v' is nil", got)
	}
}

type assertthat[T comparable] struct {
	t    *testing.T
	this T
}

func That[T comparable](t *testing.T, this T) assertthat[T] {
	return assertthat[T]{t, this}
}

func (a assertthat[T]) Equals(that T) {
	Equal(a.t, a.this, that)
}

func (a assertthat[T]) IsNil() {
	IsNil(a.t, a.this)
}

func (a assertthat[T]) IsNotNil() {
	IsNotNil(a.t, a.this)
}
