package storecore

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func FieldToSchema(in []Field) []SchemaType {
	out := make([]SchemaType, len(in))
	for k := range in {
		out[k] = in[k].DataType
	}
	return out
}

func HandleStoreSetRecord(db *Database, event *StorecoreStoreSetRecord) {
	tableName := KeyToTableName(event.TableId)

	if tableName.ResourceType != TABLE_PREFIX && tableName.ResourceType != OFFCHAINTABLE_PREFIX {
		fmt.Printf("ignoring resource: %d  %s. %s  -  %s\n", len(event.KeyTuple), tableName.ResourceType, tableName.Name, tableName.Namespace)
		return
	}

	if event.TableId == STORE_TABLES {
		tableName := KeyToTableName(event.KeyTuple[0])
		key := Key{
			world:     GetWorld(event.Raw),
			TableName: tableName,
		}

		fmt.Println("creating tables")
		fmt.Println(key)
		fmt.Println(hexutil.Encode(event.KeyTuple[0][:]))
		data := CreateNewTable(db, &key, event)
		var STORE_KEY = Key{
			world:     GetWorld(event.Raw),
			TableName: KeyToTableName(STORE_TABLES),
		}
		fmt.Println(event.StaticData)
		fmt.Println(event.DynamicData)
		fmt.Println(event.EncodedLengths)
		AddRow(db, event.KeyTuple, data, &STORE_KEY, &Raw{
			StaticData:     event.StaticData,
			DynamicData:    event.DynamicData,
			EncodedLengths: event.EncodedLengths,
		})
		return
	}

	fmt.Println("adding record to:")
	fmt.Println(tableName)
	fmt.Println("key:")

	key := Key{
		world:     GetWorld(event.Raw),
		TableName: tableName,
	}
	fmt.Println(key)
	table := db.GetTable(&key)

	if table == nil {
		fmt.Println(event.KeyTuple)
		fmt.Println(event.StaticData)
		fmt.Println(event.EncodedLengths)
		fmt.Println(event.DynamicData)
		// if tableName.Name == "FunctionSignatur" {
		// 	// This is fixes in next15
		// 	return
		// }
		fmt.Println("schema not stored")
		return
	}

	data := DecodeRecord(append(event.StaticData, append(event.EncodedLengths[:], event.DynamicData...)...), SchemaTypes{
		Static:           FieldToSchema(table.ValueSchema.StaticFields),
		Dynamic:          FieldToSchema(table.ValueSchema.DynamicFields),
		StaticDataLength: 0,
	})
	AddRow(db, event.KeyTuple, data, &key, &Raw{
		StaticData:     event.StaticData,
		DynamicData:    event.DynamicData,
		EncodedLengths: event.EncodedLengths,
	})
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
			DataType: staticType[k],
			Name:     names.Cols[k],
		}
	}

	dynamicFields := make([]Field, len(dynamicType))
	for k := range dynamicType {
		dynamicFields[k] = Field{
			DataType: dynamicType[k],
			Name:     names.Cols[k+len(staticFields)],
		}
	}

	return Schema{
		StaticFields:  staticFields,
		DynamicFields: dynamicFields,
	}
}

func CreateNewTable(db *Database, key *Key, event *StorecoreStoreSetRecord) *DecodedData {
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
		RawData:     map[string]*Raw{},
	}

	db.AddTable(key, &table)

	return data
}

func KeyTupleToDBKey(keyTuple [][32]byte) string {
	temp := []byte{}
	for k := range keyTuple {
		temp = append(temp, keyTuple[k][:]...)
	}
	return hexutil.Encode(temp)
}

func AddRow(db *Database, keyTuple [][32]byte, data *DecodedData, tableKey *Key, raw *Raw) {
	rowKey := KeyTupleToDBKey(keyTuple)
	values := []interface{}{}
	for _, v := range data.StaticData {
		values = append(values, v.Value.(string))
	}
	for _, v := range data.DynamicData {
		values = append(values, v.Value.(string))
	}
	db.AddRow(tableKey, rowKey, values, raw)
}

func HandleStoreSpliceStaticData(db *Database, event *StorecoreStoreSpliceStaticData) {
	tableName := KeyToTableName(event.TableId)
	key := Key{
		world:     GetWorld(event.Raw),
		TableName: tableName,
	}

	table := db.GetTable(&key)
	if table == nil {
		if event.TableId != RESOURCES_TABLE {
			fmt.Println("----- Table NOT found while splice static -----")
			return
		}
		fmt.Println("FIX: early table creation for resources ids")
		CreateNewTable(db, &key, &RESOURCES_CREATION_EVENT)
		table = db.GetTable(&key)
	}

	rowKey := KeyTupleToDBKey(event.KeyTuple)

	prevValue, ok := table.RawData[rowKey]
	data, _ := table.Data[rowKey]
	if !ok {
		fmt.Println("schema")
		fmt.Println(table.ValueSchema)
		prevValue = EmptyRaw(table.ValueSchema.StaticFields)
		data = make([]interface{}, len(table.ValueSchema.StaticFields)+len(table.ValueSchema.DynamicFields))
		for k := range data {
			data[k] = ""
		}
	}
	fmt.Println("lengthh------>>>>")
	fmt.Println(len(prevValue.StaticData))
	fmt.Println("start")
	fmt.Println(event.Start.Uint64())

	// We can read all the values because the Raw was init with the static length
	fmt.Println("old data")
	fmt.Println(prevValue.StaticData)
	prevValue.StaticData = append(append(prevValue.StaticData[0:event.Start.Uint64()], event.Data...), prevValue.StaticData[event.Start.Uint64()+uint64(len(event.Data)):]...)
	fmt.Println("new data")
	fmt.Println(prevValue.StaticData)

	// Decode it and store the decoded values
	decodedData := DecodeRecord(prevValue.StaticData, SchemaTypes{
		Static:           FieldToSchema(table.ValueSchema.StaticFields),
		Dynamic:          []SchemaType{},
		StaticDataLength: 0,
	})

	for k, v := range decodedData.StaticData {
		data[k] = v.Value.(string)
		fmt.Println("new value from splice: " + data[k].(string))
	}

	table.Data[rowKey] = data
}

func HandleStoreSpliceDynamicData(db *Database, event *StorecoreStoreSpliceDynamicData) {
	tableName := KeyToTableName(event.TableId)
	key := Key{
		world:     GetWorld(event.Raw),
		TableName: tableName,
	}

	table := db.GetTable(&key)
	if table == nil {
		fmt.Println("----- Table NOT found while splice dynamic -----")
		return
	}

	rowKey := KeyTupleToDBKey(event.KeyTuple)

	prevValue, ok := table.RawData[rowKey]
	data, _ := table.Data[rowKey]
	if !ok {
		fmt.Println("schema")
		fmt.Println(table.ValueSchema)
		prevValue = EmptyRaw(table.ValueSchema.StaticFields)
		data = make([]interface{}, len(table.ValueSchema.StaticFields)+len(table.ValueSchema.DynamicFields))
		for k := range data {
			data[k] = ""
		}
	}

	// We can read all the values because the Raw was init with the static length
	fmt.Println(event.Start.Uint64())
	fmt.Println(event.DeleteCount.Uint64())
	fmt.Println("old data")
	fmt.Println(prevValue.DynamicData)

	prevValue.DynamicData = append(append(prevValue.DynamicData[0:event.Start.Uint64()], event.Data...), prevValue.DynamicData[event.Start.Uint64()+event.DeleteCount.Uint64():]...)
	prevValue.EncodedLengths = event.EncodedLengths
	fmt.Println("new data")
	fmt.Println(prevValue.DynamicData)

	// Decode it and store the decoded values
	decodedData := DecodeRecord(append(prevValue.EncodedLengths[:], prevValue.DynamicData...), SchemaTypes{
		Static:           []SchemaType{},
		Dynamic:          FieldToSchema(table.ValueSchema.DynamicFields),
		StaticDataLength: 0,
	})

	for k, v := range decodedData.DynamicData {
		data[k+len(table.ValueSchema.StaticFields)] = v.Value.(string)
		fmt.Println("new value from splice: " + data[k+len(table.ValueSchema.StaticFields)].(string))
	}

	table.Data[rowKey] = data
}

func HandleStoreDeleteRecord(db *Database, event *StorecoreStoreDeleteRecord) {
	tableName := KeyToTableName(event.TableId)
	key := Key{
		world:     GetWorld(event.Raw),
		TableName: tableName,
	}
	rowKey := KeyTupleToDBKey(event.KeyTuple)
	db.RemoveRow(&key, rowKey)
}
