package iteration

import (
	"fmt"
	testinghelpers "learninggowithtests/helpers"
	"testing"
)

func TestRepeatPhrase(t *testing.T) {
	// t.Skip()
	t.Run("returns str with input repeated n times", func(t *testing.T) {
		got := RepeatPhrase("Dennis", 5)
		want := "DennisDennisDennisDennisDennis"

		testinghelpers.AssetCorrectMsg(t, got, want, "Dennis, 5")
	})
}


func BenchmarkRepeatPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatPhrase("Dennis", 5)
	}
}

func ExampleRepeatPhrase() {
	fmt.Println(RepeatPhrase("Dennis", 2))
	// Output: DennisDennis
}
