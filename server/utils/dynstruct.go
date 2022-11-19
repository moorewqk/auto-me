package utils

import (
	"fmt"
	"github.com/goldeneggg/structil/dynamicstruct"
)

type DynamicSt struct {
	Pool map[string]interface{}
}

func MakeStruct() {
	type Hoge struct {
		Key   string
		Value interface{}
	}

	hogePtr := &Hoge{
		Key:   "keystr",
		Value: "valuestr",
	}

	// Add fields using Builder with AddXXX method chain
	b := dynamicstruct.NewBuilder().
		AddString("StringField").
		AddInt("IntField").
		AddFloat32("Float32Field").
		AddBool("BoolField").
		AddMap("MapField", "name", 123).
		AddStructPtr("StructPtrField", hogePtr).
		AddSlice("SliceField", 12).
		AddInterfaceWithTag("SomeObjectField", true, `json:"some_object_field"`)

	// Remove removes a field by assigned name
	b = b.Remove("Float32Field")

	// SetStructName sets the name of DynamicStruct
	// Note: Default struct name is "DynamicStruct"
	b.SetStructName("MyStruct")

	// Build returns a DynamicStruct
	ds, err := b.Build()
	if err != nil {
		panic(err)
	}

	// Print struct definition with Definition method
	// Struct fields are automatically orderd by field name
	fmt.Println(ds.Definition())
	return
}
