package assert

import (
	"reflect"
	"testing"
)

// shamelessly lifted from matryer/is
// https://github.com/matryer/is/blob/master/is.go
func isNil(got any) bool {
	if got == nil {
		return true
	}
	value := reflect.ValueOf(got)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}

// shamelessly lifted from matryer/is
// https://github.com/matryer/is/blob/master/is.go
func areEqual(a, b any) bool {
	if isNil(a) && isNil(b) {
		return true
	}
	if isNil(a) || isNil(b) {
		return false
	}
	if reflect.DeepEqual(a, b) {
		return true
	}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	return aValue == bValue
}

type harness struct {
	t *testing.T
}

func New(t *testing.T) harness {
	return harness{t}
}

type assertthat struct {
	h    harness
	this any
}

func (h harness) That(object any) assertthat {
	return assertthat{h, object}
}

func (a assertthat) Equals(object any) {
	a.h.t.Helper()
	if !areEqual(a.this, object) {
		a.h.t.Errorf("got: (%T) %v; want: (%T) %v",
			a.this,
			a.this,
			object,
			object,
		)
	}
}

func (a assertthat) IsNil() {
	a.h.t.Helper()
	if !isNil(a.this) {
		a.h.t.Errorf("%v (%T) is not nil!", a.this, a.this)
	}
}

func (a assertthat) IsNotNil() {
	a.h.t.Helper()
	if isNil(a.this) {
		a.h.t.Errorf("%v (%T) is nil!", a.this, a.this)
	}
}

func (a assertthat) IsTrue() {
	a.h.t.Helper()
	if a.this != true {
		a.h.t.Errorf("%v (%T) is not true!", a.this, a.this)
	}
}

func (a assertthat) IsFalse() {
	a.h.t.Helper()
	if a.this != false {
		a.h.t.Errorf("%v (%T) is not false!", a.this, a.this)
	}
}
