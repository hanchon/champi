package storecore

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStaticFieldDecoder(t *testing.T) {
	tests := map[string]struct {
		data        string
		schema      SchemaType
		expectedOne interface{}
	}{
		"BYTES32": {
			data:        "0x0000000000000000000000000000000000000000000000000000000000000001",
			schema:      SchemaType(95),
			expectedOne: "0x0000000000000000000000000000000000000000000000000000000000000001",
		},

		"BOOL": {
			data:        "0x00",
			schema:      SchemaType(96),
			expectedOne: "false",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOne := DecodeStaticField(test.schema, StringToByteArray(test.data), 0)
			if diff := cmp.Diff(actualOne, test.expectedOne); diff != "" {
				t.Errorf("static data doesn't match: %v", diff)
			}

		})
	}
}

func TestDynamicFieldDecoder(t *testing.T) {
	tests := map[string]struct {
		data        string
		schema      SchemaType
		expectedOne interface{}
	}{
		"BOOL_ARRAY": {
			data:        "0x00",
			schema:      BOOL_ARRAY,
			expectedOne: []uint8(`["false"]`),
		},
		"UINT8_ARRAY": {
			data:        "0x00",
			schema:      UINT8_ARRAY,
			expectedOne: []uint8(`["0"]`),
		},
		"UINT256_ARRAY": {
			data:        "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			schema:      UINT256_ARRAY,
			expectedOne: []uint8(`["115792089237316195423570985008687907853269984665640564039457584007913129639935"]`),
		},
		// TODO: this should be negative 1
		"INT8_ARRAY": {
			data:        "0xff",
			schema:      INT8_ARRAY,
			expectedOne: []uint8(`["-1"]`),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOne := DecodeDynamicField(test.schema, StringToByteArray(test.data))
			if diff := cmp.Diff(actualOne, test.expectedOne); diff != "" {
				t.Errorf("static data doesn't match: %v", diff)
			}

		})
	}
}
