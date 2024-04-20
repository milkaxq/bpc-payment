package config

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func InitDB() {

	db, err := badger.Open(badger.DefaultOptions("badger"))
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
