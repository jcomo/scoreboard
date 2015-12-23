package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, want, got interface{}) {
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func ArrayEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
