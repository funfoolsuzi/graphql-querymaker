package querymaker

import (
	"reflect"
	"strings"
)

const (
	tab          = "  "
	fieldTagName = "graphqlvar"
)

type query struct {
	name         string
	subGqlFields []*gqlField
	variables    map[string]string
}

type gqlField struct {
	name         string
	subGqlFields []*gqlField
	variableName string
}

func newQuery(name string) *query {
	return &query{
		name:      name,
		variables: map[string]string{},
	}
}

// addSubfieldsFromStruct adds struct-subfield to query-subfield recursively
func (q *query) addSubfieldsFromStruct(f *gqlField, typeStruct reflect.Type) {
	subs := make([]*gqlField, 0, typeStruct.NumField())

	// loop thru fields
	for i := 0; i < typeStruct.NumField(); i++ {
		sf := typeStruct.Field(i)
		fieldName := getGqlFieldFromStructField(sf)
		if fieldName == "" {
			continue
		}

		sub := &gqlField{
			name: fieldName,
		}
		q.variableScan(sf, sub)
		t := getTypeOrElementType(sf.Type)
		if t.Kind() == reflect.Struct {
			q.addSubfieldsFromStruct(sub, t)
		}
		subs = append(subs, sub)
	}
	if f == nil { // means it's the root/operation
		q.subGqlFields = subs
		return
	}

	f.subGqlFields = subs
}

func (q *query) variableScan(sf reflect.StructField, f *gqlField) {
	varParts := strings.Split(sf.Tag.Get(fieldTagName), ",")
	if len(varParts) != 2 || varParts[0] == "" || varParts[1] == "" {
		return
	}
	q.variables[varParts[0]] = varParts[1]
	f.variableName = varParts[0]
	return
}

// getGqlFieldFromStructField reads data in the sf(StructField)
func getGqlFieldFromStructField(sf reflect.StructField) string {
	first := sf.Name[:1]
	if strings.ToLower(first) == first {
		return "" // unexported field should be ignored
	}

	jsonTag := sf.Tag.Get("json")
	if jsonTag == "" {
		return sf.Name
	}

	parts := strings.Split(jsonTag, ",")
	if parts[0] == "-" {
		return "" // suppressed field
	}

	return parts[0]
}
