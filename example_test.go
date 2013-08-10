package reflections_test

import (
    "fmt"
    "github.com/oleiade/reflections"
)

func ExampleGetField() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    fieldsToExtract := []string{"FirstField", "ThirdField"}

    for _, fieldName := range fieldsToExtract {
        value, err := reflections.GetField(s, fieldName)
        fmt.Println(value)
    }
}

func ExampleHasField() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    // has == true
    has, _ := reflections.HasField(s, "FirstField")

    // has == false
    has, _ := reflections.HasField(s, "FourthField")
}

func ExampleFields() {
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
}

func ExampleItems() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    var structItems map[string]interface{}

    // Items will return a field name to
    // field value map
    structItems, _ = reflections.Items(s)
}

func ExampleTags() {
    s := MyStruct {
        FirstField: "first value",      `matched:"first tag"`
        SecondField: 2,                 `matched:"second tag"`
        ThirdField: "third value",      `unmatched:"third tag"`
    }

    var structTags map[string]string

    // Tags will return a field name to tag content
    // map. Nota that only field with the tag name
    // you've provided which will be matched.
    // Here structTags will contain:
    // {
    //     "FirstField": "first tag",
    //     "SecondField": "second tag",
    // }
    structTags, _ = reflections.Tags(s, "matched")
}