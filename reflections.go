// Copyright (c) 2013 Théo Crevon
//
// See the file LICENSE for copying permission.

/*
Package reflections provides high level abstractions above the
reflect library.

Reflect library is very low-level and as can be quite complex when it comes to do simple things like accessing a structure field value, a field tag...

The purpose of reflections package is to make developers life easier when it comes to introspect structures at runtime.
It's API is freely inspired from python language (getattr, setattr, hasattr...) and provides a simplified access to structure fields and tags.
*/
package reflections

import (
	"errors"
	"fmt"
	"reflect"
)

// GetField returns the value of the provided obj field. obj param
// has to be a struct type.
func GetField(obj interface{}, name string) (interface{}, error) {
	if !isStruct(obj) {
		return nil, errors.New("Cannot use GetField on a non-struct interface")
	}

	val := reflect.ValueOf(obj)
	value := val.FieldByName(name)

	if !value.IsValid() {
		return nil, fmt.Errorf("No such field: %s in obj", name)
	}

	return value.Interface(), nil
}

// GetFieldKind returns the kind of the provided obj field. obj param
// has to be a struct type.
func GetFieldKind(obj interface{}, name string) (reflect.Kind, error) {
	if !isStruct(obj) {
		return reflect.Invalid, errors.New("Cannot use GetField on a non-struct interface")
	}

	val := reflect.ValueOf(obj)
	value := val.FieldByName(name)

	if !value.IsValid() {
		return reflect.Invalid, fmt.Errorf("No such field: %s in obj", name)
	}

	return value.Type().Kind(), nil
}

// SetField sets the provided obj field with provided value. obj param has
// to be a pointer to a struct, otherwise it will soundly fail. Provided
// value type should match with the struct field you're trying to set.
func SetField(obj interface{}, name string, value interface{}) error {
	// Fetch the field reflect.Value
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	// If obj field value is not settable an error is thrown
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

// HasField checks if the provided field name is part of a struct.
func HasField(obj interface{}, name string) (bool, error) {
	if !isStruct(obj) {
		return false, errors.New("Cannot use HasField on a non-struct interface")
	}

	structValue := reflect.TypeOf(obj)
	structField, ok := structValue.FieldByName(name)
	if !ok || !isExportableField(structField) {
		return false, nil
	}

	return true, nil
}

// Fields returns the struct fields names list
func Fields(obj interface{}) ([]string, error) {
	if !isStruct(obj) {
		return nil, errors.New("Cannot use Fields on a non-struct interface")
	}

	structType := reflect.TypeOf(obj)
	fieldsCount := structType.NumField()

	var fields []string
	for i := 0; i < fieldsCount; i++ {
		field := structType.Field(i)
		if isExportableField(field) {
			fields = append(fields, field.Name)
		}
	}

	return fields, nil
}

// Items returns the field - value struct pairs as a map
func Items(obj interface{}) (map[string]interface{}, error) {
	if !isStruct(obj) {
		return nil, errors.New("Cannot use Items on a non-struct interface")
	}

	structType := reflect.TypeOf(obj)
	structValue := reflect.ValueOf(obj)
	fieldsCount := structType.NumField()

	items := make(map[string]interface{})

	for i := 0; i < fieldsCount; i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		// Make sure only exportable and addressable fields are
		// returned by Items
		if isExportableField(field) {
			items[field.Name] = fieldValue.Interface()
		}
	}

	return items, nil
}

// Tags lists the struct tag fields
func Tags(obj interface{}, key string) (map[string]string, error) {
	if !isStruct(obj) {
		return nil, errors.New("Cannot use Tags on a non-struct interface")
	}

	structType := reflect.TypeOf(obj)
	fieldsCount := structType.NumField()

	tags := make(map[string]string)

	for i := 0; i < fieldsCount; i++ {
		structField := structType.Field(i)

		if isExportableField(structField) {
			tags[structField.Name] = structField.Tag.Get(key)
		}
	}

	return tags, nil
}

func isExportableField(field reflect.StructField) bool {
	// PkgPath is empty for exported fields.
	return field.PkgPath == ""
}

func isStruct(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Struct
}

func isPointer(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Ptr
}
