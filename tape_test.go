package main

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, close := createTempFile(t, "12345")
	defer close()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
