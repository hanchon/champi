package storecore

import (
	"fmt"
	"reflect"
	"testing"
)

type Value struct {
	data     interface{}
	dataType reflect.Type
}

func TestReflect(t *testing.T) {
	data := int8(-1)
	v := Value{
		data:     data,
		dataType: reflect.TypeOf(data),
	}

	switch reflect.ValueOf(v.data).Type().Name() {
	case "int8":
		fmt.Println(int8(-1) + v.data.(int8))
	}

}
