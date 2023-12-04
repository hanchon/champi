package storecore

func DecodeRecord(raw []byte, schema SchemaTypePair) *DecodedData {
	return DecodeData(raw, schema)
}
