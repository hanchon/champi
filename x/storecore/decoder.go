package storecore

import "strconv"

type SchemaTypes struct {
	Static           []SchemaType
	Dynamic          []SchemaType
	StaticDataLength uint64
}

func DecodeRecord(raw []byte, schemaTypePair SchemaTypes) *DecodedData {
	var bytesOffset uint64 = 0
	ret := NewDecodedData()

	// Decode static fields.
	for _, fieldType := range schemaTypePair.Static {
		value := DecodeStaticField(fieldType, raw, bytesOffset)
		bytesOffset += GetStaticByteLength(fieldType)
		ret.AddStaticElement(value, fieldType)
	}

	// Decode dynamic fields.
	if len(schemaTypePair.Dynamic) > 0 {
		dynamicDataSlice := raw[bytesOffset : bytesOffset+32]
		bytesOffset += 32

		dataLayout, _ := DecodeEncodedLengths([32]byte(dynamicDataSlice))
		// TODO: the second value must be the same as the sum of dataLayout

		for i, fieldType := range schemaTypePair.Dynamic {
			dataLength, _ := strconv.ParseInt(dataLayout[i], 10, 64)
			value := DecodeDynamicField(fieldType, raw[bytesOffset:bytesOffset+uint64(dataLength)])
			ret.AddDynamicElement(value, fieldType)

			bytesOffset += uint64(dataLength)
		}
	}

	return ret
}
