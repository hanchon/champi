package storecore

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
)

func StringToByteArray(s string) []byte {
	bytes, err := hexutil.Decode(s)
	if err != nil {
		panic("error decoding the hex string")
	}
	return bytes
}

func TestDecodeRecord(t *testing.T) {
	tests := map[string]struct {
		data        string
		schema      SchemaTypePair
		expectedOne *DecodedData
	}{
		"static only": {
			data: "0x0000000100000000000000000000000000000002000000000000000000000000000000000000000b0000000008000000000000130000000300000004736f6d6520737472696e67",
			schema: SchemaTypePair{
				Static:           []SchemaType{3, 15},
				Dynamic:          []SchemaType{101, 197},
				StaticDataLength: 20,
			},

			expectedOne: &DecodedData{
				Data: map[string]*DataSchemaTypePair{
					"STRING":       {Data: string("some string"), SchemaType: SchemaType(197)},
					"UINT128":      {Data: string("2"), SchemaType: SchemaType(15)},
					"UINT32":       {Data: string("1"), SchemaType: SchemaType(3)},
					"UINT32_ARRAY": {Data: []uint8(`["3","4"]`), SchemaType: SchemaType(101)},
				},
				Schema: []SchemaType{SchemaType(3), SchemaType(15), SchemaType(101), SchemaType(197)},
			},
		},
		"emtpy record": {
			data: "0x0000000000000000000000000000000000000000000000000000000000000000",
			schema: SchemaTypePair{
				Static:           []SchemaType{},
				Dynamic:          []SchemaType{197, 197},
				StaticDataLength: 0,
			},

			expectedOne: &DecodedData{
				Data: map[string]*DataSchemaTypePair{
					"STRING":  {Data: string(""), SchemaType: SchemaType(197)},
					"STRING2": {Data: string(""), SchemaType: SchemaType(197)},
				},
				Schema: []SchemaType{SchemaType(197), SchemaType(197)},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOne := DecodeRecord(StringToByteArray(test.data), test.schema)
			if diff := cmp.Diff(actualOne, test.expectedOne); diff != "" {
				t.Errorf("static data doesn't match: %v", diff)
			}

		})
	}
}
