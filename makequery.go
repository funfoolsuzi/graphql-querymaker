package querymaker

import (
	"reflect"
)

// MakeQuery generates GraphQL query string and add it to
func MakeQuery(query interface{}) (string, []string) {

	t := getTypeOrElementType(reflect.TypeOf(query))

	op := newQuery(t.Name())

	op.addSubfieldsFromStruct(nil, t)

	return op.generateTmplNVars()
}

func getTypeOrElementType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice:
		return getTypeOrElementType(t.Elem())
	default:
		return t
	}
}
