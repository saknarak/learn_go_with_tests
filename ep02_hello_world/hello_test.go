package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("case 1: ", func(t *testing.T) {
		got := Hello("Somchai", "en")
		want := "Hello, Somchai"
		assertCorrectName(t, got, want)
	})
	t.Run("case 2: ", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World"
		assertCorrectName(t, got, want)
	})
	t.Run("case 3: in Thai", func(t *testing.T) {
		got := Hello("Somsak", "th")
		want := "Sawasdee, Somsak"
		assertCorrectName(t, got, want)
	})
	t.Run("case 4: in Thai", func(t *testing.T) {
		got := Hello("", "th")
		want := "Sawasdee, World"
		assertCorrectName(t, got, want)
	})
	t.Run("case 5: in Japan", func(t *testing.T) {
		got := Hello("Somsri", "jp")
		want := "Ohiyo, Somsri"
		assertCorrectName(t, got, want)
	})
	t.Run("case 6: default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectName(t, got, want)
	})
}

func assertCorrectName(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}
