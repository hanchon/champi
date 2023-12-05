package storecore

import (
	"encoding/json"
	"strconv"
)

func DecodeDynamicField(schemaType SchemaType, encodingSlice []byte) interface{} {
	switch schemaType {
	case BYTES:
		return handleBytes(encodingSlice)
	case STRING:
		return handleString(encodingSlice)
	default:
		// Try to decode as an array.
		staticSchemaType := (schemaType - 98)
		if staticSchemaType > 97 {
			panic("Unknown dynamic field type:" + schemaType.String())
		}

		fieldLength := GetStaticByteLength(staticSchemaType)
		arrayLength := len(encodingSlice) / int(fieldLength)
		array := make([]interface{}, arrayLength)
		for i := 0; i < arrayLength; i++ {
			array[i] = DecodeStaticField(staticSchemaType, encodingSlice, uint64(i)*fieldLength)
		}

		arr, err := json.Marshal(array)
		if err != nil {
			panic("Could not marshal array:" + err.Error())
		}
		return arr
	}
}

func DecodeStaticField(schemaType SchemaType, encoding []byte, bytesOffset uint64) interface{} {
	if schemaType >= UINT8 && schemaType <= UINT256 {
		return handleUint(encoding[bytesOffset : bytesOffset+uint64(schemaType)+1])
	} else if schemaType >= INT8 && schemaType <= INT256 {
		return handleInt(encoding[bytesOffset:bytesOffset+uint64(schemaType-UINT256)], schemaType)
	} else if schemaType >= BYTES1 && schemaType <= BYTES32 {
		return handleBytes(encoding[bytesOffset : bytesOffset+uint64(schemaType-INT256)])
	} else if schemaType == BOOL {
		return handleBool(encoding[bytesOffset])
	} else if schemaType == ADDRESS {
		return handleAddress(encoding[bytesOffset : bytesOffset+20])
	} else if schemaType == STRING {
		return handleString(encoding[bytesOffset:])
	} else {
		panic("Unknown static field type: " + schemaType.String())
	}
}

func BytesToLength(data [32]byte) ([]string, string) {
	totalLength := DecodeStaticField(UINT56, data[32-7:], 0)
	eachLength := DecodeDynamicField(UINT40_ARRAY, data[:32-7])

	var m []string
	json.Unmarshal(eachLength.([]byte), &m)
	return reverse(m), totalLength.(string)
}

func DecodeData(encoding []byte, schemaTypePair SchemaTypePair) *DecodedData {
	var bytesOffset uint64 = 0
	ret := NewDecodedDataFromSchemaTypePair(schemaTypePair)

	// Decode static fields.
	for _, fieldType := range schemaTypePair.Static {
		value := DecodeStaticField(fieldType, encoding, bytesOffset)
		bytesOffset += GetStaticByteLength(fieldType)

		ret.Set(fieldType.String(), &DataSchemaTypePair{
			Data:       value,
			SchemaType: fieldType,
		})
	}

	// Decode dynamic fields.
	if len(schemaTypePair.Dynamic) > 0 {
		dynamicDataSlice := encoding[bytesOffset : bytesOffset+32]
		bytesOffset += 32

		dataLayout, _ := BytesToLength([32]byte(dynamicDataSlice))
		// TODO: the second value must be the same as the sum of dataLayout

		for i, fieldType := range schemaTypePair.Dynamic {
			dataLength, _ := strconv.ParseInt(dataLayout[i], 10, 64)
			value := DecodeDynamicField(fieldType, encoding[bytesOffset:bytesOffset+uint64(dataLength)])
			bytesOffset += uint64(dataLength)

			ret.Set(fieldType.String(), &DataSchemaTypePair{
				Data:       value,
				SchemaType: fieldType,
			})
		}
	}

	return ret
}
