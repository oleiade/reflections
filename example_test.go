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