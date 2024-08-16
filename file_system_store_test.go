package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	data := `[
	  {"Name": "Cleo", "Wins": 10},
	  {"Name": "Chris", "Wins": 33}
	]`

	t.Run("leage from a reader", func(t *testing.T) {
		database, close := createTempFile(t, data)
		defer close()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player store", func(t *testing.T) {
		database, close := createTempFile(t, data)
		defer close()

		store := NewFileSystemPlayerStore(database)

		got, _ := store.GetPlayerScore("Chris")
		assertScoreEquals(t, got, 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, close := createTempFile(t, data)
		defer close()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Cleo")

		got, _ := store.GetPlayerScore("Cleo")
		assertScoreEquals(t, got, 11)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, close := createTempFile(t, data)
		defer close()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got, _ := store.GetPlayerScore("Pepper")
		assertScoreEquals(t, got, 1)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %q", err)
	}

	_, err = tmpfile.Write([]byte(initialData))
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
