package storecore

import (
	"encoding/json"
)

func DecodeDynamicField(schemaType SchemaType, raw []byte) string {
	switch schemaType {
	case BYTES:
		return handleBytes(raw)
	case STRING:
		return handleString(raw)
	default:
		// Try to decode as an array.
		staticSchemaType := (schemaType - 98)
		if staticSchemaType > 97 {
			panic("Unknown dynamic field type:" + schemaType.String())
		}

		fieldLength := GetStaticByteLength(staticSchemaType)
		arrayLength := len(raw) / int(fieldLength)
		array := make([]interface{}, arrayLength)
		for i := 0; i < arrayLength; i++ {
			array[i] = DecodeStaticField(staticSchemaType, raw, uint64(i)*fieldLength)
		}

		arr, err := json.Marshal(array)
		if err != nil {
			panic("Could not marshal array:" + err.Error())
		}
		return string(arr)
	}
}

func DecodeStaticField(schemaType SchemaType, raw []byte, bytesOffset uint64) interface{} {
	if schemaType >= UINT8 && schemaType <= UINT256 {
		return handleUint(raw[bytesOffset : bytesOffset+uint64(schemaType)+1])
	} else if schemaType >= INT8 && schemaType <= INT256 {
		return handleInt(raw[bytesOffset:bytesOffset+uint64(schemaType-UINT256)], schemaType)
	} else if schemaType >= BYTES1 && schemaType <= BYTES32 {
		return handleBytes(raw[bytesOffset : bytesOffset+uint64(schemaType-INT256)])
	} else if schemaType == BOOL {
		return handleBool(raw[bytesOffset])
	} else if schemaType == ADDRESS {
		return handleAddress(raw[bytesOffset : bytesOffset+20])
	} else if schemaType == STRING {
		return handleString(raw[bytesOffset:])
	} else {
		panic("Unknown static field type: " + schemaType.String())
	}
}

func DecodeEncodedLengths(data [32]byte) ([]string, string) {
	totalLength := DecodeStaticField(UINT56, data[32-7:], 0)
	eachLength := DecodeDynamicField(UINT40_ARRAY, data[:32-7])

	var m []string
	json.Unmarshal([]byte(eachLength), &m)
	return reverse(m), totalLength.(string)
}
