package maps

import (
	"testing"
)

func TestDict(t *testing.T) {
	myDict := Dict{
		Words: map[string]string{"dennis": "b mithamo"},
	}

	t.Run("Search - word exists", func(t *testing.T) {
		got, err := myDict.Search("dennis")
		want := "b mithamo"

		if err != nil {
			t.Errorf("got err %q want nil", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Search - word not exists", func(t *testing.T) {
		got, err := myDict.Search("deniece")
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if err == nil {
			t.Error("want error got nil")
		}

		if err != NOT_FOUND_ERR {
			t.Errorf("want err %q got err %q", NOT_FOUND_ERR, err)
		}
	})

}
