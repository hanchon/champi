package storecore

import (
	"encoding/json"
	"fmt"
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

func TestDynamicFieldDecoder(t *testing.T) {
	tests := map[string]struct {
		data        string
		schema      SchemaType
		expectedOne interface{}
	}{
		"BOOL_ARRAY": {
			data:        "0x00",
			schema:      BOOL_ARRAY,
			expectedOne: []uint8("[false]"),
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
				var m []int8
				json.Unmarshal(actualOne.([]byte), &m)
				fmt.Println("m")
				fmt.Println(m)
				fmt.Println(actualOne)
				fmt.Println(test.expectedOne)
				// q := binary.BigEndian.Uint16([]byte("\xff"))
				// fmt.Println("q")
				// fmt.Println(q)
				// fmt.Println(uint8(q))
				// fmt.Println(string(actualOne.([]uint8)))

				// big, _ := hexutil.DecodeBig("0xff")
				// fmt.Println(big)
				// fmt.Println(int8(big.Int64()))

				// a, err := strconv.ParseInt("0xff", 0, 16)
				// fmt.Println(err)
				// fmt.Println(uint8(a))
				// temp, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(string(actualOne.([]uint8)), "[\"", ""), "\"]", ""), 10, 8)
				// fmt.Println(temp)
				t.Errorf("static data doesn't match: %v", diff)
			}

		})
	}
}
