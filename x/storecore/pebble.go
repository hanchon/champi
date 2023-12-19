package storecore

import (
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
)

type PebbleDB struct {
	DB *pebble.DB
}

func CreateDB(name string) *PebbleDB {
	db, err := pebble.Open(name, &pebble.Options{
		FormatMajorVersion: pebble.FormatVirtualSSTables,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &PebbleDB{DB: db}
}

func (db *PebbleDB) Close() error {
	return db.DB.Close()

}

func (db *PebbleDB) AddSchema() {

}

func (db *PebbleDB) Set(key []byte, value []byte) error {
	return db.DB.Set(key, value, pebble.Sync)
}

func (db *PebbleDB) Get(key []byte) ([]byte, error) {
	value, closer, err := db.DB.Get(key)
	if err != nil {
		if closer != nil {
			closer.Close()
		}
		return value, err
	}

	if errCloser := closer.Close(); errCloser != nil {
		return value, errCloser
	}
	return value, nil
}

func (db PebbleDB) qwe() {
	key := []byte("hello")
	if err := db.DB.Set(key, []byte("world"), pebble.Sync); err != nil {
		log.Fatal(err)
	}
	value, closer, err := db.DB.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", key, value)
	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
}
