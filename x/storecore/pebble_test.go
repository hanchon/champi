package storecore

import (
	"os"
	"testing"

	"github.com/cockroachdb/pebble"
	"github.com/google/go-cmp/cmp"
)

func TestPebbleDB(t *testing.T) {
	os.RemoveAll("/tmp/test.db")

	db := CreateDB("/tmp/test.db")
	defer db.Close()

	key := []byte("testkey")
	value := []byte("value")

	_, err := db.Get(key)
	if err != pebble.ErrNotFound {
		t.Errorf("error key should not exist")
	}

	if err := db.Set(key, value); err != nil {
		t.Errorf("error setting the key")
	}

	dbValue, err := db.Get(key)
	if diff := cmp.Diff(err, nil); diff != "" {
		t.Errorf("error should be nil")
	}

	if diff := cmp.Diff(dbValue, value); diff != "" {
		t.Errorf("value should be the same as the one stored: %s", diff)
	}

}
