// Copyright (c) 2013 Théo Crevon
//
// See the file LICENSE for copying permission.

/*
Package reflections provides high level abstractions above the
reflect library.

It's purpose is to make developers life easier when it comes to introspect structs. It's API is freely inspired
from python language (getattr, setattr, hasattr...) and provides an enhanced access to struct fields and tags.
*/

package reflections

import (
    "fmt"
    "errors"
    "unsafe"
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
        errMsg := fmt.Sprintf("No such field: %s in obj", name)
        return nil, errors.New(errMsg)
    }

    return value.Interface(), nil
}

// SetField sets the provided obj field with provided value. obj param has
// to be a pointer to a struct, otherwise it will soundly fail. Provided
// value type should match with the struct field you're trying to set.
func SetField(obj interface{}, name string, value interface{}) error {
    // Fetch the field reflect.Value
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        errMsg := fmt.Sprintf("No such field: %s in obj", name)
        return errors.New(errMsg)
    }

    // If obj field value is not settable an error is thrown
    if !structFieldValue.CanSet() {
        errMsg := fmt.Sprintf("Cannot set %s field value", name)
        return errors.New(errMsg)
    }

    invalidTypeError := errors.New("Provided value type didn't match obj field type")

    switch value.(type) {
    case bool:
        if structFieldValue.Type().Kind() != reflect.Bool {
            return invalidTypeError
        }
        structFieldValue.SetBool(value.(bool))
    case int:
        if structFieldValue.Type().Kind() != reflect.Int64 {
            return invalidTypeError
        }
        structFieldValue.SetInt(value.(int64))
    case uint64:
        if structFieldValue.Type().Kind() != reflect.Uint64 {
            return invalidTypeError
        }
        structFieldValue.SetUint(value.(uint64))
    case float64:
        if structFieldValue.Type().Kind() != reflect.Float64 {
            return invalidTypeError
        }
        structFieldValue.SetFloat(value.(float64))
    case complex128:
        if structFieldValue.Type().Kind() != reflect.Complex128 {
            return invalidTypeError
        }
        structFieldValue.SetComplex(value.(complex128))
    case string:
        if structFieldValue.Type().Kind() != reflect.String {
            return invalidTypeError
        }
        structFieldValue.SetString(value.(string))
    case []byte:
        if structFieldValue.Type().Kind() != reflect.Slice {
            return invalidTypeError
        }
        structFieldValue.SetBytes(value.([]byte))
    case unsafe.Pointer:
        if structFieldValue.Type().Kind() != reflect.Ptr {
            return invalidTypeError
        }
        structFieldValue.SetPointer(value.(unsafe.Pointer))
    default:
        return errors.New("Unknow field type")
    }

    return nil
}

// HasField checks if the provided field name is part of a struct.
func HasField(obj interface{}, name string) (bool, error) {
    if !isStruct(obj) {
        return false, errors.New("Cannot use HasField on a non-struct interface")
    }

    structValue := reflect.TypeOf(obj)
    structField, ok := structValue.FieldByName(name)
    if !ok || !isExportableField(&structField) {
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
        if isExportableField(&field) {
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
        if isExportableField(&field) {
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

        if isExportableField(&structField) {
            tags[structField.Name] = structField.Tag.Get(key)
        }
    }

    return tags, nil
}

func isExportableField(field *reflect.StructField) bool {
    // golang variables must start with a letter,
    // so checking if first letter is within [a-z]
    // is sufficient here
    if field.Name[0] >= 97 && field.Name[0] <= 122 {
        return false
    }
    return true
}

func isStruct(obj interface{}) bool {
    return reflect.TypeOf(obj).Kind() == reflect.Struct
}

func isPointer(obj interface{}) bool {
    return reflect.TypeOf(obj).Kind() == reflect.Ptr
}
