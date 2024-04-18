package config

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

var Config struct {
	DB *badger.DB
}

func InitDB() {
	opt := badger.DefaultOptions("").WithInMemory(true)

	db, err := badger.Open(opt)
	if err != nil {
		log.Fatal(err)
	}

	Config.DB = db
}
