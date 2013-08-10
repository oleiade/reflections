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