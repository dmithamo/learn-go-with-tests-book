package maps

import (
	"testing"
)

func TestDict(t *testing.T) {
	myDict := Dict{
		Words: map[string]string{"dennis": "b mithamo"},
	}

	t.Run("Search word exists", func(t *testing.T) {
		got, err := myDict.Search("dennis")
		want := "b mithamo"

		if err != nil {
			t.Errorf("got err %q want nil", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Search word not exists", func(t *testing.T) {
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

	t.Run("Add new word", func(t *testing.T) {
		word := "b"
		defn := "dennis mithamo"

		got := myDict.Add(word, defn)

		if got != nil {
			t.Errorf("want nil err got %q", got)
		}

		// search to verify addn
		currentDef, err := myDict.Search("b")
		if currentDef != defn || err != nil {
			t.Errorf("word %q with defn %q not added", word, defn)
		}
	})

	t.Run("Add word already exists", func(t *testing.T) {
		word := "b"
		defn := "dennis mithamo"

		// Add first
		myDict.Add(word, defn)

		// Retry Add
		err := myDict.Add(word, defn)
		if err == nil {
			t.Error("expected error did not get error")
		}

		if err != WORD_EXISTS_ERR {
			t.Errorf("expected err %q got err %q", WORD_EXISTS_ERR, err)
		}
	})

}
