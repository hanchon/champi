package storecore

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFieldDecoder(t *testing.T) {
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
			expectedOne: false,
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
