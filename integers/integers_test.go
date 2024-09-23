package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("handles multiple  numbers", func(t *testing.T) {
		got := Adder(2, 2, 2)
		expected := 6
		assetCorrectMsg(t, got, expected)
	})
}

func assetCorrectMsg(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdder() {
	fmt.Println(Adder(2, 2))
	// Output 6
}
