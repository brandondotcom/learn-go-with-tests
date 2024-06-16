package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Brandon")
	want := "Hello Brandon"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
