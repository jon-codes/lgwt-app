package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatalf("problem opening %q: %v", dbFileName, err)
	}

	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000: %v", err)
	}
}
