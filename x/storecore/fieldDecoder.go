package storecore

import (
	"encoding/json"
	"fmt"
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
			// return ""
		}

		// Allocate an array of the correct size.
		fieldLength := GetStaticByteLength(staticSchemaType)
		arrayLength := len(encodingSlice) / int(fieldLength)
		array := make([]interface{}, arrayLength)
		// Iterate and decode each element as a static field.
		for i := 0; i < arrayLength; i++ {
			array[i] = DecodeStaticField(staticSchemaType, encodingSlice, uint64(i)*fieldLength)
		}

		arr, err := json.Marshal(array)
		if err != nil {
			panic("Could not marshal array:" + err.Error())
			// return ""
		}

		return arr
	}
}

func DecodeStaticField(schemaType SchemaType, encoding []byte, bytesOffset uint64) interface{} {
	// To avoid a ton of duplicate handling code per each schema type, we handle
	// using ranges, since the schema types are sequential in specific ranges.

	// UINT8 - UINT256 is the first range. We add one to the schema type to get the
	// number of bytes to read, since enums start from 0 and UINT8 is the first one.
	if schemaType >= UINT8 && schemaType <= UINT256 {
		return handleUint(encoding[bytesOffset : bytesOffset+uint64(schemaType)+1])
	} else
	// INT8 - INT256 is the second range. We subtract UINT256 from the schema type
	// to account for the first range and re-set the bytes count to start from 1.
	if schemaType >= INT8 && schemaType <= INT256 {
		return handleInt(encoding[bytesOffset : bytesOffset+uint64(schemaType-UINT256)])
	} else
	// BYTES is the third range. We subtract INT256 from the schema type to account
	// for the previous ranges and re-set the bytes count to start from 1.
	if schemaType >= BYTES1 && schemaType <= BYTES32 {
		return handleBytes(encoding[bytesOffset : bytesOffset+uint64(schemaType-INT256)])
	} else
	// BOOL is a standalone schema type.
	if schemaType == BOOL {
		return handleBool(encoding[bytesOffset])
	} else
	// ADDRESS is a standalone schema type.
	if schemaType == ADDRESS {
		return handleAddress(encoding[bytesOffset : bytesOffset+20])
	} else
	// STRING is a standalone schema type.
	if schemaType == STRING {
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

	// Where the decoded data is stored.
	data := NewDecodedDataFromSchemaTypePair(schemaTypePair)

	// Decode static fields.
	for _, fieldType := range schemaTypePair.Static {
		value := DecodeStaticField(fieldType, encoding, bytesOffset)
		bytesOffset += GetStaticByteLength(fieldType)

		// Save a mapping of FIELD TYPE (string) -> (FIELD VALUE (interface{}), FIELD TYPE (SchemaType))
		data.Set(fieldType.String(), &DataSchemaTypePair{
			Data:       value,
			SchemaType: fieldType,
		})
	}

	// Decode dynamic fields.
	if len(schemaTypePair.Dynamic) > 0 {
		// dynamicDataSlice := encoding[schemaTypePair.StaticDataLength : schemaTypePair.StaticDataLength+32]
		dynamicDataSlice := encoding[bytesOffset : bytesOffset+32]
		// dynamicDataSlice := encoding[schemaTypePair.StaticDataLength:]
		fmt.Println("len(dynamicDataSlice)")
		fmt.Println(len(dynamicDataSlice))
		bytesOffset += 32

		temp := DecodeStaticField(UINT56, dynamicDataSlice[32-7:], 0)
		fmt.Println("total")
		fmt.Println(temp)

		temp2 := DecodeDynamicField(UINT40_ARRAY, dynamicDataSlice[:32-7])
		fmt.Println("len")
		fmt.Println(temp2)
		var m []string
		json.Unmarshal(temp2.([]byte), &m)
		fmt.Println(m)
		dataLayout := reverse(m)

		temp3 := uint64(0)
		for _, v := range temp2.([]uint8) {
			temp3 += uint64(v)
		}
		fmt.Println(temp3)

		for i, fieldType := range schemaTypePair.Dynamic {
			fmt.Println(fieldType)
			fmt.Println("qweqwe")
			fmt.Println("datalayout")
			fmt.Println(i)
			fmt.Println(dataLayout[i])

			// temp := DecodeStaticField(UINT56, bytes.Trim(dynamicDataSlice[32-7:], "\x00"), 0)

			// offset := 4 + i*2
			// dataLength := new(big.Int).SetBytes(dynamicDataSlice[offset : offset+2]).Uint64()
			dataLength, _ := strconv.ParseInt(dataLayout[i], 10, 64)
			fmt.Println(dataLength)
			value := DecodeDynamicField(fieldType, encoding[bytesOffset:bytesOffset+uint64(dataLength)])
			bytesOffset += uint64(dataLength)

			// Save a mapping of FIELD TYPE (string) -> (FIELD VALUE (interface{}), FIELD TYPE (SchemaType))
			data.Set(fieldType.String(), &DataSchemaTypePair{
				Data:       value,
				SchemaType: fieldType,
			})

		}
	}

	return data
}
