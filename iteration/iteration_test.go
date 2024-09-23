package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatPhrase(t *testing.T) {
	t.Skip()
	t.Run("returns str with input repeated n times", func(t *testing.T) {
		got := RepeatPhrase("Dennis", 5)
		want := "DennisDennisDennisDennisDennis"

		assetCorrectMsg(t, got, want)
	})
}

func assetCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
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
