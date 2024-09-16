package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("empty lang defaults to 'Hiya,'", func(t *testing.T) {
		got := Hello("", "Dennis")
		want := "Hiya, Dennis!"

		assetCorrectMsg(t, got, want)
	})

	t.Run("empty name defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hiya, world!"

		assetCorrectMsg(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("SPA", "Khalegi")
		want := "Hola, Khalegi!"

		assetCorrectMsg(t, got, want)
	})
}

func assetCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
