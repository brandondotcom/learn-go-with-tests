package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// yuck
func captureStdout() (r *os.File, w *os.File, stdout *os.File) {
	stdout = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	return
}

func readAndRestoreStdout(r, w, stdout *os.File) string {
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdout

	return string(out)
}

func assertCorrectOutput(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	t.Run("with defaults when no options specified", func(t *testing.T) {
		r, w, stdout := captureStdout()

		Greet()

		got := readAndRestoreStdout(r, w, stdout)
		want := "Hello, World"

		assertCorrectOutput(t, got, want)
	})

	t.Run("with name", func(t *testing.T) {
		r, w, stdout := captureStdout()

		Greet(WithName("Brandon"))

		got := readAndRestoreStdout(r, w, stdout)
		want := "Hello, Brandon"

		assertCorrectOutput(t, got, want)
	})

	t.Run("with writer", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(WithWriter(&buffer))

		got := buffer.String()
		want := "Hello, World"

		assertCorrectOutput(t, got, want)
	})

	t.Run("with all options", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(
			WithName("Chris"),
			WithWriter(&buffer),
		)

		got := buffer.String()
		want := "Hello, Chris"

		assertCorrectOutput(t, got, want)
	})
}
