package storecore

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type SchemaTypeKV struct {
	Key   *SchemaTypePair `json:"key"`
	Value *SchemaTypePair `json:"value"`
}

func (pair *SchemaTypeKV) Flatten() []SchemaType {
	return append(pair.Key.Flatten(), pair.Value.Flatten()...)
}

type SchemaTypePair struct {
	Static           []SchemaType `json:"static"`
	Dynamic          []SchemaType `json:"dynamic"`
	StaticDataLength uint64       `json:"static_data_length"`
}

func (tuple *SchemaTypePair) Flatten() []SchemaType {
	return append(tuple.Static, tuple.Dynamic...)
}

func SchemaTypeKVFromPairs(key *SchemaTypePair, value *SchemaTypePair) *SchemaTypeKV {
	return &SchemaTypeKV{
		Key:   key,
		Value: value,
	}
}

func CombineSchemaTypePair(schemaTypePair SchemaTypePair) []SchemaType {
	return schemaTypePair.Flatten()
}

func StringifySchemaTypes(schemaType []SchemaType) []string {
	var types []string
	for _, t := range schemaType {
		types = append(types, strings.ToLower(t.String()))
	}
	return types
}

func CombineStringifySchemaTypes(schemaType []SchemaType) string {
	return strings.Join(StringifySchemaTypes(schemaType), ",")
}

type DataSchemaTypePair struct {
	Data       interface{}
	SchemaType SchemaType
}

type DecodedData struct {
	Data   []DataSchemaTypePair
	Schema []SchemaType
}

// NewDecodedDataFromSchemaTypePair creates a new instance of DecodedData with the provided SchemaTypePair.
//
// Parameters:
// - schemaTypePair (SchemaTypePair): The SchemaTypePair to use for the DecodedData instance.
//
// Returns:
// (*DecodedData): The new DecodedData instance.
func NewDecodedDataFromSchemaTypePair(schemaTypePair SchemaTypePair) *DecodedData {
	return &DecodedData{
		Data:   []DataSchemaTypePair{},
		Schema: CombineSchemaTypePair(schemaTypePair),
	}
}

// NewDecodedDataFromSchemaType creates a new instance of DecodedData with the provided list of SchemaType.
//
// Parameters:
// - schemaType ([]SchemaType): The list of SchemaType to use for the DecodedData instance.
//
// Returns:
// (*DecodedData): The new DecodedData instance.
func NewDecodedDataFromSchemaType(schemaType []SchemaType) *DecodedData {
	return &DecodedData{
		Data:   []DataSchemaTypePair{},
		Schema: schemaType,
	}
}

// Length returns the length of the schema types in the DecodedData instance.
//
// Returns:
// (int): The length of the schema types in the DecodedData instance.
func (d *DecodedData) Length() int {
	return len(d.Schema)
}

// Set sets the value for the given key in the DecodedData instance.
//
// Parameters:
// - key (string): The key to set the value for.
// - value (*DataSchemaTypePair): The value to set for the key.
//
// Returns:
// - void.
func (d *DecodedData) Set(key string, value *DataSchemaTypePair) {
	d.Data = append(d.Data, *value)
}

// SchemaTypes retrieves a slice of all the schema types in the DecodedData instance.
//
// Returns:
// ([]SchemaType): A slice of all the schema types in the DecodedData instance.
func (d *DecodedData) SchemaTypes() []SchemaType {
	return d.Schema
}

func handleBytes(encoding []byte) string {
	// Hex-encode bytes for legibility.
	return hexutil.Encode(encoding)
}

func handleUint(encoding []byte) string {
	return new(big.Int).SetBytes(encoding).String()
}

func handleInt(encoding []byte, schemaType SchemaType) string {
	value := new(big.Int).SetBytes(encoding)
	pow := int64(((schemaType - 32 + 1) * 8) - 1)
	compareValue := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(pow), big.NewInt(0))
	if value.Cmp(compareValue) == 1 {
		offset := compareValue.Mul(compareValue, big.NewInt(-2))
		return value.Add(value, offset).String()
	}
	return value.String()
}

func handleBool(encoding byte) bool {
	return encoding == 1
}

func handleAddress(encoding []byte) string {
	return common.BytesToAddress(encoding).String()
}

func handleString(encoding []byte) string {
	return string(encoding)
}

func GetStaticByteLength(schemaType SchemaType) uint64 {
	if schemaType < 32 {
		// uint8-256
		return uint64(schemaType) + 1
	} else if schemaType < 64 {
		// int8-256, offset by 32
		return uint64(schemaType) + 1 - 32
	} else if schemaType < 96 {
		// bytes1-32, offset by 64
		return uint64(schemaType) + 1 - 64
	}

	// Other static types
	if schemaType == BOOL {
		return 1
	} else if schemaType == ADDRESS {
		return 20
	}

	// Return 0 for all dynamic types
	return 0
}

// SchemaTypeToSolidityType converts the specified SchemaType instance to the corresponding Solidity type.
//
// Parameters:
// - schemaType (SchemaType): The SchemaType instance to convert to a Solidity type.
//
// Returns:
// (string) - A string representing the Solidity type for the specified SchemaType instance.
func SchemaTypeToSolidityType(schemaType SchemaType) string {
	if strings.Contains(schemaType.String(), "ARRAY") {
		_type := strings.Split(schemaType.String(), "_")[0]
		return fmt.Sprintf("%s[]", strings.ToLower(_type))
	} else {
		return strings.ToLower(schemaType.String())
	}
}

// SchemaTypeToPostgresType converts the specified SchemaType instance to the corresponding PostgreSQL type.
// The function returns a string representing the PostgreSQL type for the specified SchemaType instance.
//
// Parameters:
// - schemaType (SchemaType): The SchemaType instance to convert to a PostgreSQL type.
//
// Returns:
// (string) - A string representing the PostgreSQL type for the specified SchemaType instance.
func SchemaTypeToPostgresType(schemaType SchemaType) string {
	if (schemaType >= UINT8 && schemaType <= UINT32) || (schemaType >= INT8 && schemaType <= INT32) {
		// Integer.
		return "integer"
	} else if (schemaType >= UINT64 && schemaType <= UINT256) || (schemaType >= INT64 && schemaType <= INT256) {
		// Big integer.
		return "text"
	} else if (schemaType >= BYTES1 && schemaType <= BYTES32) || (schemaType == BYTES) {
		// Bytes.
		return "text"
	} else if schemaType == BOOL {
		// Boolean.
		return "boolean"
	} else if schemaType == ADDRESS {
		// Address.
		return "text"
	} else if schemaType == STRING {
		// String.
		return "text"
	} else if (schemaType >= UINT8_ARRAY && schemaType <= UINT32_ARRAY) || (schemaType >= INT8_ARRAY && schemaType <= INT32_ARRAY) {
		// Integer array.
		return "jsonb"
	} else if (schemaType >= UINT64_ARRAY && schemaType <= UINT256_ARRAY) || (schemaType >= INT64_ARRAY && schemaType <= INT256_ARRAY) {
		// Big integer array.
		return "jsonb"
	} else if schemaType >= BYTES1_ARRAY && schemaType <= BYTES32_ARRAY {
		// Bytes array.
		return "jsonb"
	} else if schemaType == BOOL_ARRAY {
		// Boolean array.
		return "jsonb"
	} else if schemaType == ADDRESS_ARRAY {
		// Address array.
		return "jsonb"
	} else {
		// Default to text.
		return "text"
	}
}
