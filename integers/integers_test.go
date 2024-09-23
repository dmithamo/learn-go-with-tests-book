package integers

import (
	"fmt"
	testinghelpers "learninggowithtests/helpers"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("handles multiple  numbers", func(t *testing.T) {
		got := Adder(2, 2, 2)
		want := 6
		testinghelpers.AssetCorrectMsg(t, got, want, "2, 2, 2")
	})
}

func ExampleAdder() {
	fmt.Println(Adder(2, 2))
	// Output: 4
}
