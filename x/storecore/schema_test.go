package storecore

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
)

func StringTo32Byte(s string) [32]byte {
	bytes, err := hexutil.Decode(s)
	if err != nil {
		panic("error decoding the hex string")
	}
	if len(bytes) != 32 {
		panic("invalid lenght")
	}
	return [32]byte(bytes)
}

func TestGenerateSchema(t *testing.T) {
	tests := map[string]struct {
		data        string
		expectedOne []string
		expectedTwo []string
	}{
		"static only": {
			data:        "0x0001010060000000000000000000000000000000000000000000000000000000",
			expectedOne: []string{"BOOL"},
			expectedTwo: []string{},
		},
		"static only two": {
			data:        "0x00570800616100030700001f0000000000000000000000000000000000000000",
			expectedOne: []string{"ADDRESS", "ADDRESS", "UINT8", "UINT32", "UINT64", "UINT8", "UINT8", "UINT256"},
			expectedTwo: []string{},
		},
		"both values": {
			data:        "0x0001010160c20000000000000000000000000000000000000000000000000000",
			expectedOne: []string{"BOOL"},
			expectedTwo: []string{"BOOL_ARRAY"},
		},
		"both values two": {
			data:        "0x002402045f2381c3c4c500000000000000000000000000000000000000000000",
			expectedOne: []string{"BYTES32", "INT32"},
			expectedTwo: []string{"UINT256_ARRAY", "ADDRESS_ARRAY", "BYTES", "STRING"},
		},
		"register tables": {
			data:        "0x006003025f5f5fc4c40000000000000000000000000000000000000000000000",
			expectedOne: []string{"BYTES32", "BYTES32", "BYTES32"},
			expectedTwo: []string{"BYTES", "BYTES"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOne, actualTwo := GenerateSchema(StringTo32Byte(test.data))
			if diff := cmp.Diff(actualOne, test.expectedOne); diff != "" {
				t.Errorf("static data doesn't match: %v", diff)
			}

			if diff := cmp.Diff(actualTwo, test.expectedTwo); diff != "" {
				t.Errorf("static data doesn't match: %v", diff)
			}
		})
	}
}
