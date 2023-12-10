package storecore

import "fmt"

type Key struct {
	world string
	TableName
}

type Database struct {
	Tables map[Key]*Table
}

type Field struct {
	dataType SchemaType
	name     string
}

type Schema struct {
	staticFields  []Field
	dynamicFields []Field
}

type Raw struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}

func EmptyRaw(staticFields []Field) *Raw {
	length := uint64(0)
	for k := range staticFields {
		fmt.Println(staticFields[k])
		length += GetStaticByteLength(staticFields[k].dataType)
	}

	fmt.Println("empty row length")
	fmt.Println(length)

	return &Raw{
		StaticData:     make([]byte, length),
		EncodedLengths: [32]byte{},
		DynamicData:    []byte{},
	}
}

type Table struct {
	Name        Key
	KeySchema   Schema
	ValueSchema Schema
	Data        map[string][]interface{}
	RawData     map[string]*Raw
}

func NewDatabase() *Database {
	return &Database{
		Tables: map[Key]*Table{},
	}
}

func (db *Database) AddTable(key *Key, table *Table) bool {
	_, ok := db.Tables[*key]
	if ok {
		return false
	}
	db.Tables[*key] = table
	return true
}

func (db *Database) AddRow(key *Key, rowID string, values []interface{}, raw *Raw) bool {
	_, ok := db.Tables[*key]
	if !ok {
		return false
	}
	db.Tables[*key].Data[rowID] = values
	db.Tables[*key].RawData[rowID] = raw
	return true
}

func (db *Database) GetTable(key *Key) *Table {
	_, ok := db.Tables[*key]
	if !ok {
		return nil
	}
	return db.Tables[*key]
}
