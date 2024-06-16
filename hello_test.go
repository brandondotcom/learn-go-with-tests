package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Goodbye world"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
