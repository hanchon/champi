package storecore

type DataElement struct {
	Value      interface{}
	SchemaType SchemaType
}

type DecodedData struct {
	StaticData  []DataElement
	DynamicData []DataElement
}

func NewDecodedData() *DecodedData {
	return &DecodedData{
		StaticData:  []DataElement{},
		DynamicData: []DataElement{},
	}
}

func (data *DecodedData) AddStaticElement(value interface{}, schemaType SchemaType) {
	data.StaticData = append(data.StaticData, DataElement{
		Value:      value,
		SchemaType: schemaType,
	})
}

func (data *DecodedData) AddDynamicElement(value interface{}, schemaType SchemaType) {
	data.DynamicData = append(data.DynamicData, DataElement{
		Value:      value,
		SchemaType: schemaType,
	})
}
