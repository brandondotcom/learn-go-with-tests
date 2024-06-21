package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("say Hello to Brandon", func(t *testing.T) {
        got := Hello("Brandon", "")
        want := "Hello, Brandon"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say Hello to world by default", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, world"

        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"

        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Nicolas", "French")
        want := "Bonjour, Nicolas"

        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q but wanted %q", got, want)
    }
}
