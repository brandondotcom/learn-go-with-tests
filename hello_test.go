package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("say Hello to Brandon", func(t *testing.T) {
        got := Hello("Brandon")
        want := "Hello Brandon"

        if got != want {
            t.Errorf("got %q but wanted %q", got, want)
        }
    })

    t.Run("say Hello to world by default", func(t *testing.T) {
        got := Hello("")
        want := "Hello world"

        if got != want {
            t.Errorf("got %q but wanted %q", got, want)
        }
    })
}
