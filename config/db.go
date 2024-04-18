package config

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func InitDB() {
	opt := badger.DefaultOptions("").WithInMemory(true)

	db, err := badger.Open(opt)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
