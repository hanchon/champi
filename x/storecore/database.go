package storecore

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

type Table struct {
	Name        Key
	KeySchema   Schema
	ValueSchema Schema
	Data        map[string][]interface{}
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

func (db *Database) AddRow(key *Key, rowID string, values []interface{}) bool {
	_, ok := db.Tables[*key]
	if !ok {
		return false
	}
	db.Tables[*key].Data[rowID] = values
	return true
}
