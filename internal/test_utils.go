package internal

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, cond bool) {
	t.Helper()
	if !cond {
		t.Fatalf("%v is false", cond)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
}

func assertDeepEqual(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v != %v", a, b)
	}
}
