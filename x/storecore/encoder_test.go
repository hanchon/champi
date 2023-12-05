package storecore

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Log struct {
	tableId        string
	keyTuple       []string
	staticData     string
	encodedLengths string
	dynamicData    string
}

// Resources
const (
	table         = "tb"
	offchainTable = "ot"
	namespace     = "ns"
	module        = "md"
	system        = "sy"
)

//     const resourceTypeId = hexToString(sliceHex(hex, 0, 2)).replace(, "");
// const type = getResourceType(resourceTypeId);
// const namespace = hexToString(sliceHex(hex, 2, 16)).replace(/\0+$/, "");
// const name = hexToString(sliceHex(hex, 16, 32)).replace(/\0+$/, "");

// (property) valueSchema: {
const fieldLayout = "BYTES32"
const keySchema = "BYTES32"
const valueSchema = "BYTES32"
const abiEncodedKeyNames = "BYTES"
const abiEncodedFieldNames = "BYTES"

// }

// func Process(static, encodedLen, dynamic []byte) {
// 	staticDataLength := new(big.Int).SetBytes(static[0:2]).Uint64()
// 	fmt.Println(staticDataLength)
// 	numStaticFields := new(big.Int).SetBytes(static[2:3]).Uint64()
// 	fmt.Println(numStaticFields)
// 	numDynamicFields := new(big.Int).SetBytes(static[3:4]).Uint64()
// 	fmt.Println(numDynamicFields)
//
// 	staticFields := []string{}
// 	var i uint64
// 	for i = 4; i < 4+numStaticFields; i++ {
// 		fmt.Println(i)
// 		fmt.Println(SchemaType(static[i]))
// 		staticFields = append(staticFields, SchemaType(static[i]).String())
// 	}
// 	dynamicFields := []string{}
// 	for i = 4 + numStaticFields; i < 4+numStaticFields+numDynamicFields; i++ {
// 		fmt.Println(i)
// 		fmt.Println(SchemaType(static[i]))
// 		dynamicFields = append(dynamicFields, SchemaType(static[i]).String())
// 	}
//
// 	fmt.Println(staticFields)
// 	fmt.Println(dynamicFields)
//
// 	// i := 0
// 	// a := SchemaType(static[i])
// 	// fmt.Println(a)
// 	// fmt.Println(a.String())
//
// 	return
// }

func TestSetStoreRecord(t *testing.T) {
	a := Log{
		tableId:        "0x74626d756473746f72650000000000005461626c657300000000000000000000",
		keyTuple:       []string{"0x74626d756473746f72650000000000005461626c657300000000000000000000"},
		staticData:     "0x0060030220202000000000000000000000000000000000000000000000000000002001005f000000000000000000000000000000000000000000000000000000006003025f5f5fc4c40000000000000000000000000000000000000000000000",
		encodedLengths: "0x000000000000000000000000000000000000022000000000a0000000000002c0",
		dynamicData:    "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000077461626c654964000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000000b6669656c644c61796f757400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000096b6579536368656d610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b76616c7565536368656d610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012616269456e636f6465644b65794e616d657300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014616269456e636f6465644669656c644e616d6573000000000000000000000000",
	}

	// First key in keytuple is the tableID
	tableID := a.keyTuple[0]
	fmt.Println(tableID)
	tableIDBytes, _ := hexutil.Decode(tableID)
	fmt.Println(tableIDBytes)
	key := KeyToTableName([32]byte(tableIDBytes))
	fmt.Println(key)
	fmt.Println(key.Name == "Tables")

	static, _ := hexutil.Decode(a.staticData)
	encodedLen, _ := hexutil.Decode(a.encodedLengths)
	// dynamic, _ := hexutil.Decode(a.dynamicData)

	// ValuesSchema
	staticType, dynamicType := GenerateSchema([32]byte(static[len(static)-32:]))
	fmt.Println("Values Schema:")
	fmt.Println(staticType)
	fmt.Println(dynamicType)
	// KeySchema
	staticTypeKey, dynamicTypeKey := GenerateSchema([32]byte(static[len(static)-64 : len(static)-32]))
	fmt.Println("Key Schema:")
	fmt.Println(staticTypeKey)
	fmt.Println(dynamicTypeKey)
	fmt.Println("Leftovers: (Field Layout-Not used in the indexer")
	fmt.Println(static[:len(static)-64])
	fmt.Printf("%s\n", bytes.Trim(static[:len(static)-64], "\x00"))

	dynamicLen, totalLen := DecodeEncodedLengths([32]byte(encodedLen))
	fmt.Println("encoded len")
	fmt.Println(dynamicLen)
	fmt.Println(totalLen)
	// t.Fatalf("error")

	// "address": "0x3Aa5ebB10DC797CAC828524e59A333d0A371443c",
	// "keySchema": {
	//   "tableId": "bytes32",
	// },
	// "name": "Tables",
	// "namespace": "mudstore",
	// "tableId": "0x74626d756473746f72650000000000005461626c657300000000000000000000",
	// "valueSchema": {
	//   "abiEncodedFieldNames": "bytes",
	//   "abiEncodedKeyNames": "bytes",
	//   "fieldLayout": "bytes32",
	//   "keySchema": "bytes32",
	//   "valueSchema": "bytes32",
	// },

}
