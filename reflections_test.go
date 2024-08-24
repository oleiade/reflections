// Copyright (c) 2013 Théo Crevon
//
// See the file LICENSE for copying permission.

package reflections

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Dummy      string `test:"dummytag"`
	unexported uint64
	Yummy      int `test:"yummytag"`
}

func TestGetField_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	value, err := GetField(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.Equal(t, "test", value)
}

func TestGetField_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
	}

	value, err := GetField(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.Equal(t, "test", value)
}

func TestGetField_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetField(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetField_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	_, err := GetField(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetField_unexported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}

	assert.Panics(t, func() {
		GetField(dummyStruct, "unexported") //nolint:errcheck,gosec
	})
}

func TestGetFieldKind_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	kind, err := GetFieldKind(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.Equal(t, reflect.String, kind)

	kind, err = GetFieldKind(dummyStruct, "Yummy")
	require.NoError(t, err)
	assert.Equal(t, reflect.Int, kind)
}

func TestGetFieldKind_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	kind, err := GetFieldKind(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.Equal(t, reflect.String, kind)

	kind, err = GetFieldKind(dummyStruct, "Yummy")
	require.NoError(t, err)
	assert.Equal(t, reflect.Int, kind)
}

func TestGetFieldKind_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetFieldKind(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetFieldKind_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	_, err := GetFieldKind(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetFieldType_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	typeString, err := GetFieldType(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.Equal(t, "string", typeString)

	typeString, err = GetFieldType(dummyStruct, "Yummy")
	require.NoError(t, err)
	assert.Equal(t, "int", typeString)
}

func TestGetFieldType_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	typeString, err := GetFieldType(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.Equal(t, "string", typeString)

	typeString, err = GetFieldType(dummyStruct, "Yummy")
	require.NoError(t, err)
	assert.Equal(t, "int", typeString)
}

func TestGetFieldType_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetFieldType(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetFieldType_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	_, err := GetFieldType(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetFieldTag_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{}

	tag, err := GetFieldTag(dummyStruct, "Dummy", "test")
	require.NoError(t, err)
	assert.Equal(t, "dummytag", tag)

	tag, err = GetFieldTag(dummyStruct, "Yummy", "test")
	require.NoError(t, err)
	assert.Equal(t, "yummytag", tag)
}

func TestGetFieldTag_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{}

	tag, err := GetFieldTag(dummyStruct, "Dummy", "test")
	require.NoError(t, err)
	assert.Equal(t, "dummytag", tag)

	tag, err = GetFieldTag(dummyStruct, "Yummy", "test")
	require.NoError(t, err)
	assert.Equal(t, "yummytag", tag)
}

func TestGetFieldTag_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetFieldTag(dummy, "Dummy", "test")
	assert.Error(t, err)
}

func TestGetFieldTag_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{}

	_, err := GetFieldTag(dummyStruct, "obladioblada", "test")
	assert.Error(t, err)
}

func TestGetFieldTag_unexported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}

	_, err := GetFieldTag(dummyStruct, "unexported", "test")
	assert.Error(t, err)
}

func TestSetField_on_struct_with_valid_value_type(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := SetField(&dummyStruct, "Dummy", "abc")
	require.NoError(t, err)
	assert.Equal(t, "abc", dummyStruct.Dummy)
}

func TestSetField_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := SetField(&dummyStruct, "obladioblada", "life goes on")
	assert.Error(t, err)
}

func TestSetField_invalid_value_type(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := SetField(&dummyStruct, "Yummy", "123")
	assert.Error(t, err)
}

func TestSetField_non_exported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	assert.Error(t, SetField(&dummyStruct, "unexported", "fail, bitch"))
}

func TestFields_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	fields, err := Fields(dummyStruct)
	require.NoError(t, err)
	assert.Equal(t, []string{"Dummy", "Yummy"}, fields)
}

func TestFields_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	fields, err := Fields(dummyStruct)
	require.NoError(t, err)
	assert.Equal(t, []string{"Dummy", "Yummy"}, fields)
}

func TestFields_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := Fields(dummy)
	assert.Error(t, err)
}

func TestFields_with_non_exported_fields(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 6789,
		Dummy:      "test",
		Yummy:      123,
	}

	fields, err := Fields(dummyStruct)
	require.NoError(t, err)
	assert.Equal(t, []string{"Dummy", "Yummy"}, fields)
}

func TestHasField_on_struct_with_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := HasField(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.True(t, has)
}

func TestHasField_on_struct_pointer_with_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := HasField(dummyStruct, "Dummy")
	require.NoError(t, err)
	assert.True(t, has)
}

func TestHasField_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := HasField(dummyStruct, "Test")
	require.NoError(t, err)
	assert.False(t, has)
}

func TestHasField_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := HasField(dummy, "Test")
	assert.Error(t, err)
}

func TestHasField_unexported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 7890,
		Dummy:      "test",
		Yummy:      123,
	}

	has, err := HasField(dummyStruct, "unexported")
	require.NoError(t, err)
	assert.False(t, has)
}

func TestTags_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Tags(dummyStruct, "test")
	require.NoError(t, err)
	assert.Equal(t, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	}, tags)
}

func TestTags_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Tags(dummyStruct, "test")
	require.NoError(t, err)
	assert.Equal(t, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	}, tags)
}

func TestTags_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := Tags(dummy, "test")
	assert.Error(t, err)
}

func TestItems_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Items(dummyStruct)
	require.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	}, tags)
}

func TestItems_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Items(dummyStruct)
	require.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	}, tags)
}

func TestItems_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := Items(dummy)
	assert.Error(t, err)
}

//nolint:unused
func TestItems_deep(t *testing.T) {
	t.Parallel()

	type Address struct {
		Street string `tag:"be"`
		Number int    `tag:"bi"`
	}

	type unexportedStruct struct{}

	type Person struct {
		Name string `tag:"bu"`
		Address
		unexportedStruct
	}

	p := Person{}
	p.Name = "John"
	p.Street = "Decumanus maximus"
	p.Number = 17

	items, err := Items(p)
	require.NoError(t, err)
	itemsDeep, err := ItemsDeep(p)
	require.NoError(t, err)

	assert.Len(t, items, 2)
	assert.Len(t, itemsDeep, 3)
	assert.Equal(t, "John", itemsDeep["Name"])
	assert.Equal(t, "Decumanus maximus", itemsDeep["Street"])
	assert.Equal(t, 17, itemsDeep["Number"])
}

func TestGetFieldNameByTagValue(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "dummy",
		Yummy: 123,
	}

	tagJSON := "dummytag"
	field, err := GetFieldNameByTagValue(dummyStruct, "test", tagJSON)

	require.NoError(t, err)
	assert.Equal(t, "Dummy", field)
}

func TestGetFieldNameByTagValue_on_non_existing_tag(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "dummy",
		Yummy: 123,
	}

	// non existing tag value with an existing tag key
	tagJSON := "tag"
	_, errTagValue := GetFieldNameByTagValue(dummyStruct, "test", tagJSON)
	require.Error(t, errTagValue)

	// non existing tag key with an existing tag value
	tagJSON = "dummytag"
	_, errTagKey := GetFieldNameByTagValue(dummyStruct, "json", tagJSON)
	require.Error(t, errTagKey)

	// non existing tag key and value
	tagJSON = "tag"
	_, errTagKeyValue := GetFieldNameByTagValue(dummyStruct, "json", tagJSON)
	require.Error(t, errTagKeyValue)
}

//nolint:unused
func TestTags_deep(t *testing.T) {
	t.Parallel()

	type Address struct {
		Street string `tag:"be"`
		Number int    `tag:"bi"`
	}

	type unexportedStruct struct{}

	type Person struct {
		Name string `tag:"bu"`
		Address
		unexportedStruct
	}

	p := Person{}
	p.Name = "John"
	p.Street = "Decumanus maximus"
	p.Number = 17

	tags, err := Tags(p, "tag")
	require.NoError(t, err)
	tagsDeep, err := TagsDeep(p, "tag")
	require.NoError(t, err)

	assert.Len(t, tags, 2)
	assert.Len(t, tagsDeep, 3)
	assert.Equal(t, "bu", tagsDeep["Name"])
	assert.Equal(t, "be", tagsDeep["Street"])
	assert.Equal(t, "bi", tagsDeep["Number"])
}

//nolint:unused
func TestFields_deep(t *testing.T) {
	t.Parallel()

	type Address struct {
		Street string `tag:"be"`
		Number int    `tag:"bi"`
	}

	type unexportedStruct struct{}

	type Person struct {
		Name string `tag:"bu"`
		Address
		unexportedStruct
	}

	p := Person{}
	p.Name = "John"
	p.Street = "street?"
	p.Number = 17

	fields, err := Fields(p)
	require.NoError(t, err)
	fieldsDeep, err := FieldsDeep(p)
	require.NoError(t, err)

	assert.Len(t, fields, 2)
	assert.Len(t, fieldsDeep, 3)
	assert.Equal(t, "Name", fieldsDeep[0])
	assert.Equal(t, "Street", fieldsDeep[1])
	assert.Equal(t, "Number", fieldsDeep[2])
}

type SingleString string

type StringList []string

type Bar struct {
	A StringList
}

func TestAssignable(t *testing.T) {
	t.Parallel()

	var b Bar
	expected := []string{"a", "b", "c"}
	require.NoError(t, SetField(&b, "A", expected))
	assert.Equal(t, StringList(expected), b.A)

	err := SetField(&b, "A", []int{0, 1, 2})
	require.Error(t, err)
	assert.Equal(t, "provided value type not assignable to obj field type",
		err.Error())
}
