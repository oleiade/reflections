# Reflections

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Build Status](https://github.com/oleiade/reflections/actions/workflows/go.yml/badge.svg)](https://github.com/oleiade/reflections/actions/workflows/go.yml)
[![Go Documentation](https://pkg.go.dev/badge/github.com/oleiade/reflections)](https://pkg.go.dev/github.com/oleiade/reflections)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleiade/reflections)](https://goreportcard.com/report/github.com/oleiade/reflections)
![Go Version](https://img.shields.io/github/go-mod/go-version/oleiade/reflections)

The `reflections` library provides high-level abstractions on top of the go language standard `reflect` library.

My experience of the `reflect` library's API is that it's somewhat low-level and unintuitive. Using it can rapidly become pretty complex, daunting, and scary, especially when doing simple things like accessing a structure field value, a field tag, etc.

The `reflections` package aims to make developers' life easier when it comes to introspect struct values at runtime. Its API is inspired by the python language `getattr,` `setattr,` and `hasattr` set of methods and provides simplified access to structure fields and tags.

## Documentation

Head to the [documentation](https://pkg.go.dev/github.com/oleiade/reflections) to get more details on the library's API.

## Usage

## Accessing structure fields

### GetField

`GetField` returns the content of a structure field. For example, it can be beneficial when you want to iterate over struct-specific field values. You can provide `GetField` a structure or a pointer to a struct as the first argument.

```go
 s := MyStruct {
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 fieldsToExtract := []string{"FirstField", "ThirdField"}

 for _, fieldName := range fieldsToExtract {
 value, err := reflections.GetField(s, fieldName)
 DoWhatEverWithThatValue(value)
 }
```

### GetFieldKind

`GetFieldKind` returns the [`reflect.Kind`](http://golang.org/src/pkg/reflect/type.go?s=6916:6930#L189) of a structure field. You can use it to operate type assertion over a structure field at runtime. You can provide `GetFieldKind` a structure or a pointer to structure as the first argument.

```go
 s := MyStruct{
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 var firstFieldKind reflect.String
 var secondFieldKind reflect.Int
 var err error

 firstFieldKind, err = GetFieldKind(s, "FirstField")
 if err != nil {
 log.Fatal(err)
 }

 secondFieldKind, err = GetFieldKind(s, "SecondField")
 if err != nil {
 log.Fatal(err)
 }
```

### GetFieldType

`GetFieldType` returns the string literal of a structure field type. You can use it to operate type assertion over a structure field at runtime. You can provide `GetFieldType` a structure or a pointer to structure as the first argument.

```go
 s := MyStruct{
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 var firstFieldKind string
 var secondFieldKind string
 var err error

 firstFieldKind, err = GetFieldType(s, "FirstField")
 if err != nil {
 log.Fatal(err)
 }

 secondFieldKind, err = GetFieldType(s, "SecondField")
 if err != nil {
 log.Fatal(err)
 }
```

### GetFieldTag

`GetFieldTag` extracts a specific structure field tag. You can provide `GetFieldTag` a structure or a pointer to structure as the first argument.

```go
 s := MyStruct{}

 tag, err := reflections.GetFieldTag(s, "FirstField", "matched")
 if err != nil {
 log.Fatal(err)
 }
 fmt.Println(tag)

 tag, err = reflections.GetFieldTag(s, "ThirdField", "unmatched")
 if err != nil {
 log.Fatal(err)
 }
 fmt.Println(tag)
```

### HasField

`HasField` asserts a field exists through the structure. You can provide `HasField` a struct or a pointer to a struct as the first argument.

```go
 s := MyStruct {
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 // has == true
 has, _ := reflections.HasField(s, "FirstField")

 // has == false
 has, _ := reflections.HasField(s, "FourthField")
```

### Fields

`Fields` returns the list of structure field names so that you can access or modify them later. You can provide `Fields` with a struct or a pointer to a struct as the first argument.

```go
 s := MyStruct {
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 var fields []string

 // Fields will list every structure exportable fields.
 // Here, it's content would be equal to:
 // []string{"FirstField", "SecondField", "ThirdField"}
 fields, _ = reflections.Fields(s)
```

### Items

`Items` returns the structure's field name to the values map. You can provide `Items` with a struct or a pointer to structure as the first argument.

```go
 s := MyStruct {
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 var structItems map[string]interface{}

 // Items will return a field name to
 // field value map
 structItems, _ = reflections.Items(s)
```

### Tags

`Tags` returns the structure's fields tag with the provided key. You can provide `Tags` with a struct or a pointer to a struct as the first argument.

```go
 s := MyStruct {
 FirstField: "first value", `matched:"first tag"`
 SecondField: 2, `matched:"second tag"`
 ThirdField: "third value", `unmatched:"third tag"`
 }

 var structTags map[string]string

 // Tags will return a field name to tag content
 // map. N.B that only field with the tag name
 // you've provided will be matched.
 // Here structTags will contain:
 // {
 // "FirstField": "first tag",
 // "SecondField": "second tag",
 // }
 structTags, _ = reflections.Tags(s, "matched")
```

### Set a structure field value

`SetField` update's a structure's field value with the one provided. Note that you can't set unexported fields and that the field and value types must match.

```go
 s := MyStruct {
 FirstField: "first value",
 SecondField: 2,
 ThirdField: "third value",
 }

 //To be able to set the structure's values,
 // it must be passed as a pointer.
 _ := reflections.SetField(&s, "FirstField", "new value")

 // If you try to set a field's value using the wrong type,
 // an error will be returned
 err := reflection.SetField(&s, "FirstField", 123) // err != nil
```

## Important notes

- **unexported fields** cannot be accessed nor set using the `reflections` library. The Go lang standard `reflect` library intentionally prohibits unexported fields values access or modifications.

## Contribute

- Check for open issues or open a new issue to start a discussion around a feature idea or a bug.
- Fork `the repository`\_ on GitHub to start making your changes to the **master** branch (or branch off of it).
- Write tests showing that the bug was fixed or the feature works as expected.
- Send a pull request and bug the maintainer until it gets merged and published. :) Make sure to add yourself to AUTHORS\_.

[the repository](http://github.com/oleiade/reflections)
[AUTHORS](https://github.com/oleiade/reflections/blob/master/AUTHORS.md)
