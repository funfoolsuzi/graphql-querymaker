package querymaker

import (
	"reflect"
)

func MakeQuery(source interface{}) (string, error) {

	t := getTypeOrElementTypeOf(source)

	op := newQuery(t.Name())

	op.addSubfieldsFromStruct(nil, t)

	return op.String(), nil
}

func getTypeOrElementTypeOf(source interface{}) reflect.Type {
	t := reflect.TypeOf(source)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func getTypeOrElementType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice:
		return getTypeOrElementType(t.Elem())
	default:
		return t
	}
}
