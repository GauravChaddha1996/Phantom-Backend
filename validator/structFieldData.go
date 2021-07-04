package validator

import "reflect"

type StructFieldData struct {
	Name  string
	Type  reflect.Type
	Kind  reflect.Kind
	Tag   reflect.StructTag
	Elem  reflect.Type
	Value reflect.Value
}

func makeStructFieldData(field reflect.StructField, value reflect.Value) StructFieldData {
	fieldType := field.Type
	fieldKind := fieldType.Kind()
	var fieldElem reflect.Type
	if fieldKind == reflect.Ptr {
		fieldElem = fieldType.Elem()
	}
	return StructFieldData{
		Name:  field.Name,
		Type:  field.Type,
		Kind:  field.Type.Kind(),
		Tag:   field.Tag,
		Elem:  fieldElem,
		Value: value,
	}
}
