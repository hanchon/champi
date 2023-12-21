package storecore

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
)

//	type Key struct {
//		world string
//		TableName
//	}
//
//	type Database struct {
//		Tables map[Key]*Table
//	}
type PebbleField struct {
	DataType SchemaType `json:"dt"`
	Name     string     `json:"name"`
}

type PebbleSchema struct {
	StaticFields  []PebbleField `json:"sf"`
	DynamicFields []PebbleField `json:"df"`
}

//
// type Raw struct {
// 	StaticData     []byte   `json:"sd"`
// 	EncodedLengths [32]byte `json:"el"`
// 	DynamicData    []byte   `json:"dd"`
// }

// type Table struct {
// 	Name        Key                      `json:"key"`
// 	KeySchema   Schema                   `json:"ks"`
// 	ValueSchema Schema                   `json:"vs"`
// 	Data        map[string][]interface{} `json:"d"`
// 	RawData     map[string]*Raw          `json:"r"`
// }

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

func SchemaKey(world string, key []byte) []byte {
	return append(append([]byte("schema"), []byte(world)...), key...)
}

func RowValuesKey(world string, table string, key []byte) []byte {
	return append(append(append([]byte("row"), []byte(world)...), []byte(table)...), key...)
}

func RowRawKey(world string, table string, key []byte) []byte {
	return append(append(append([]byte("raw"), []byte(world)...), []byte(table)...), key...)
}

func (db *PebbleDB) AddSchema(world string, key []byte, schema PebbleSchema) error {
	schemaKey := SchemaKey(world, key)
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return err
	}
	return db.Set(schemaKey, schemaBytes)
}

func (db *PebbleDB) GetSchema(world string, key []byte) (*PebbleSchema, error) {
	var schema PebbleSchema
	bytes, err := db.Get(SchemaKey(world, key))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &schema); err != nil {
		return nil, err
	}
	return &schema, nil
}

func (db *PebbleDB) AddRow(world string, table string, key []byte, values []string, raw Raw) error {
	rowValuesBytes, err := json.Marshal(values)
	if err != nil {
		return err
	}

	rowRawBytes, err := json.Marshal(raw)
	if err != nil {
		return err
	}

	db.Set(RowValuesKey(world, table, key), rowValuesBytes)
	db.Set(RowRawKey(world, table, key), rowRawBytes)

	return nil
}

func (db *PebbleDB) GetRowValues(world string, table string, key []byte) ([]string, error) {
	bytes, err := db.Get(RowValuesKey(world, table, key))
	if err != nil {
		return []string{}, err
	}

	var values []string
	if err := json.Unmarshal(bytes, &values); err != nil {
		return []string{}, err
	}

	return values, nil
}

func (db *PebbleDB) GetRowRaw(world string, table string, key []byte) (Raw, error) {
	bytes, err := db.Get(RowRawKey(world, table, key))
	if err != nil {
		return Raw{}, err
	}

	var raw Raw
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return Raw{}, err
	}
	return raw, nil
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

func (db *PebbleDB) SetValues(key []byte, values [][]byte) error {

	return nil
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
