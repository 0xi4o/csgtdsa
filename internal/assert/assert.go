package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

// func ComareSlices[T comparable](t *testing.T, actual, expected T) {
// 	t.Helper()
//
// 	if !slices.Equal(actual, expected) {
// 		t.Errorf("expected %v but got %v", expected, actual)
// 	}
// }
