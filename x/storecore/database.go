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
	DataType SchemaType `json:"dt"`
	Name     string     `json:"name"`
}

type Schema struct {
	StaticFields  []Field `json:"sf"`
	DynamicFields []Field `json:"df"`
}

type Raw struct {
	StaticData     []byte   `json:"sd"`
	EncodedLengths [32]byte `json:"el"`
	DynamicData    []byte   `json:"dd"`
}

type Table struct {
	Name        Key                      `json:"key"`
	KeySchema   Schema                   `json:"ks"`
	ValueSchema Schema                   `json:"vs"`
	Data        map[string][]interface{} `json:"d"`
	RawData     map[string]*Raw          `json:"r"`
}

func EmptyRaw(staticFields []Field) *Raw {
	length := uint64(0)
	for k := range staticFields {
		fmt.Println(staticFields[k])
		length += GetStaticByteLength(staticFields[k].DataType)
	}

	fmt.Println("empty row length")
	fmt.Println(length)

	return &Raw{
		StaticData:     make([]byte, length),
		EncodedLengths: [32]byte{},
		DynamicData:    []byte{},
	}
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

func (db *Database) RemoveRow(key *Key, rowID string) bool {
	_, ok := db.Tables[*key]
	if !ok {
		return false
	}
	delete(db.Tables[*key].Data, rowID)
	delete(db.Tables[*key].RawData, rowID)
	return true
}

func (db *Database) GetTable(key *Key) *Table {
	_, ok := db.Tables[*key]
	if !ok {
		return nil
	}
	return db.Tables[*key]
}
