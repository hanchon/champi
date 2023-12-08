package storecore

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// "tbstoreTables"
var STORE_TABLES = [32]byte{116, 98, 115, 116, 111, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 84, 97, 98, 108, 101, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var STORE_TABLES_KEY = TableName{
	ResourceType: "tb",
	Namespace:    "store",
	Name:         "Tables",
}

func HandleStoreSetRecord(db *Database, event *StorecoreStoreSetRecord) {
	fmt.Println("using world:")
	fmt.Println(GetWorld(event.Raw))

	tableName := KeyToTableName(event.KeyTuple[0])
	key := Key{
		world:     GetWorld(event.Raw),
		TableName: tableName,
	}

	if event.TableId == STORE_TABLES {
		fmt.Println("creating tables")
		fmt.Println(key)
		table, data := CreateNewTable(db, &key, event)
		fmt.Println("1")
		fmt.Println(table)
		AddRow(db, &key, event, data, table)
		fmt.Println("2")
		fmt.Println(table)
		// CreateTable(db, &key, event)
	}

}

func AddNamesToSchema(staticType, dynamicType []SchemaType, data DataElement) Schema {
	x, err := hexutil.Decode(data.Value.(string))
	if err != nil {
		fmt.Println("invalid data:" + err.Error())
		return Schema{}

	}
	names, err := DecodeNames(x)
	if err != nil {
		fmt.Println("invalid field names")
		return Schema{}
	}

	if len(names.Cols) != len(staticType)+len(dynamicType) {
		fmt.Println("error invalid len on names")
		return Schema{}
	}

	staticFields := make([]Field, len(staticType))
	for k := range staticType {
		staticFields[k] = Field{
			dataType: staticType[k],
			name:     names.Cols[k],
		}
	}

	dynamicFields := make([]Field, len(dynamicType))
	for k := range dynamicType {
		dynamicFields[k] = Field{
			dataType: dynamicType[k],
			name:     names.Cols[k+len(staticFields)],
		}
	}

	return Schema{
		staticFields:  staticFields,
		dynamicFields: dynamicFields,
	}
}

func CreateNewTable(db *Database, key *Key, event *StorecoreStoreSetRecord) (*Table, *DecodedData) {
	// KeySchema
	keyStatic, keyDynamic := GenerateSchema([32]byte(event.StaticData[len(event.StaticData)-64 : len(event.StaticData)-32]))
	// fmt.Println("Key Schema:")
	// fmt.Println(keyStatic)
	// fmt.Println(keyDynamic)

	// ValuesSchema
	valuesStatic, valuesDynamic := GenerateSchema([32]byte(event.StaticData[len(event.StaticData)-32:]))
	// fmt.Println("Values Schema:")
	// fmt.Println(valuesStatic)
	// fmt.Println(valuesDynamic)

	// Field Names
	data := DecodeRecord(append(event.StaticData, append(event.EncodedLengths[:], event.DynamicData...)...), SchemaTypes{
		// Hardcoded values based on MUD doc for encoding tables
		Static:           []SchemaType{BYTES32, BYTES32, BYTES32},
		Dynamic:          []SchemaType{BYTES, BYTES},
		StaticDataLength: GetStaticByteLength(BYTES32) * 3,
	})

	table := Table{
		Name:        *key,
		KeySchema:   AddNamesToSchema(keyStatic, keyDynamic, data.DynamicData[0]),
		ValueSchema: AddNamesToSchema(valuesStatic, valuesDynamic, data.DynamicData[1]),
		Data:        map[string][]interface{}{},
	}

	db.AddTable(key, &table)

	return &table, data
}

func AddRow(db *Database, key *Key, event *StorecoreStoreSetRecord, data *DecodedData, table *Table) {
	temp := []byte{}
	for k := range event.KeyTuple {
		temp = append(temp, event.KeyTuple[k][:]...)
	}

	rowKey := hexutil.Encode(temp)

	values := []interface{}{}
	for _, v := range data.StaticData {
		values = append(values, v.Value.(string))
	}
	for _, v := range data.DynamicData {
		values = append(values, v.Value.(string))
	}
	table.Data[rowKey] = values
}

func CreateTable(db *Database, key *Key, event *StorecoreStoreSetRecord) {
	// Table creation
	if event.TableId == STORE_TABLES {
		fmt.Print("Registering table:")
		fmt.Println(KeyToTableName(event.KeyTuple[0]))

		// ValuesSchema
		staticTypes, dynamicTypes := GenerateSchema([32]byte(event.StaticData[len(event.StaticData)-32:]))
		fmt.Println("Values Schema:")
		fmt.Println(staticTypes)
		fmt.Println(dynamicTypes)
		// KeySchema
		staticTypeKey, dynamicTypeKey := GenerateSchema([32]byte(event.StaticData[len(event.StaticData)-64 : len(event.StaticData)-32]))
		fmt.Println("Key Schema:")
		fmt.Println(staticTypeKey)
		fmt.Println(dynamicTypeKey)
		// Field Names
		data := DecodeRecord(append(event.StaticData, append(event.EncodedLengths[:], event.DynamicData...)...), SchemaTypes{
			Static: []SchemaType{BYTES32, BYTES32, BYTES32},
			Dynamic: []SchemaType{
				BYTES,
				BYTES,
			},
			StaticDataLength: GetStaticByteLength(BYTES32) * 3,
		})

		x, _ := hexutil.Decode(data.DynamicData[0].Value.(string))
		keyNames, err := DecodeNames(x)
		if err != nil {
			fmt.Println("invalid field names")
		}
		fmt.Println("Key Names:")
		fmt.Println(keyNames.Cols)

		x, _ = hexutil.Decode(data.DynamicData[1].Value.(string))
		valueNames, err := DecodeNames(x)
		if err != nil {
			fmt.Println("invalid field names")
		}
		fmt.Println("Values Names:")
		fmt.Println(valueNames.Cols)

		staticFields := []Field{}
		for k := range staticTypes {
			staticFields = append(staticFields, Field{
				dataType: staticTypes[k],
				name:     valueNames.Cols[k],
			})
		}

		dynamicFields := []Field{}
		for k := range dynamicTypes {
			dynamicFields = append(dynamicFields, Field{
				dataType: dynamicTypes[k],
				name:     valueNames.Cols[k+len(staticFields)],
			})
		}
		rowKey := hexutil.Encode(event.KeyTuple[0][:])

		values := []interface{}{}
		for _, v := range data.StaticData {
			values = append(values, v.Value.(string))
		}
		for _, v := range data.DynamicData {
			values = append(values, v.Value.(string))
		}

		table := Table{
			// KeySchema: Schema{
			// 	staticFields:  staticTypeKey,
			// 	dynamicFields: dynamicTypeKey,
			// },
			// ValueSchema: Schema{
			// 	staticFields:  staticFields,
			// 	dynamicFields: dynamicFields,
			// },
			Data: map[string][]interface{}{
				rowKey: values,
			},
		}
		db.Tables[*key] = &table
		fmt.Println(db)
	}
}
