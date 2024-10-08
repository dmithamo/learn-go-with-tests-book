package testinghelpers

import (
	"reflect"
	"testing"
)

// AssetCorrectMsg is a reusable fn for checking whether a test passes
func AssetCorrectMsg(t testing.TB, got, want, given interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given %v", got, want, given)
	}
}
