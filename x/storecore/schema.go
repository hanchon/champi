package storecore

import (
	"math/big"

	"github.com/umbracle/ethgo/abi"
)

// The last 32 bytes from the static field are the schema
func GenerateSchema(data [32]byte) (staticFields, dynamicFields []string) {
	staticFields = []string{}
	dynamicFields = []string{}

	numStaticFields := new(big.Int).SetBytes(data[2:3]).Uint64()
	numDynamicFields := new(big.Int).SetBytes(data[3:4]).Uint64()

	var i uint64
	for i = 4; i < 4+numStaticFields; i++ {
		staticFields = append(staticFields, SchemaType(data[i]).String())
	}
	for i = 4 + numStaticFields; i < 4+numStaticFields+numDynamicFields; i++ {
		dynamicFields = append(dynamicFields, SchemaType(data[i]).String())
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
