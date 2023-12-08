package storecore

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/umbracle/ethgo/abi"
)

// The last 32 bytes from the static field are the schema
func GenerateSchema(data [32]byte) (staticFields, dynamicFields []SchemaType) {
	staticFields = []SchemaType{}
	dynamicFields = []SchemaType{}

	// staticDataLength := new(big.Int).SetBytes(data[0:2]).Uint64()
	numStaticFields := new(big.Int).SetBytes(data[2:3]).Uint64()
	numDynamicFields := new(big.Int).SetBytes(data[3:4]).Uint64()

	var i uint64
	for i = 4; i < 4+numStaticFields; i++ {
		staticFields = append(staticFields, SchemaType(data[i]))
	}

	// TODO: validate that staticFields length is equal to the value sent inside the message
	for i = 4 + numStaticFields; i < 4+numStaticFields+numDynamicFields; i++ {
		dynamicFields = append(dynamicFields, SchemaType(data[i]))
	}

	return staticFields, dynamicFields
}

type SchemaNames struct {
	Cols []string
}

func DecodeNames(data []byte) (SchemaNames, error) {
	_type := abi.MustNewType("tuple(string[] cols)")
	outStruct := SchemaNames{
		Cols: []string{},
	}
	err := _type.DecodeStruct(data, &outStruct)
	return outStruct, err
}

type TableName struct {
	ResourceType string
	Namespace    string
	Name         string
}

func (t *TableName) Equals(other *TableName) bool {
	return t.ResourceType == other.ResourceType && t.Namespace == other.Namespace && t.Name == other.Name
}

func KeyToTableName(key [32]byte) TableName {
	return TableName{
		ResourceType: fmt.Sprintf("%s", bytes.Trim(key[0:2], "\x00")),
		Namespace:    fmt.Sprintf("%s", bytes.Trim(key[2:16], "\x00")),
		Name:         fmt.Sprintf("%s", bytes.Trim(key[16:32], "\x00")),
	}
}
