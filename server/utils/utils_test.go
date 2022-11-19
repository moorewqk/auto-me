package utils

import (
	"fmt"
	"testing"
)

type a struct{}
type b struct{}
type c struct{}

func TestObject(t *testing.T) {
	{
		a := "2"
		b := "2s"
		fmt.Println(a, b)
	}

	//var (
	//	 stA = make([]interface{},0)
	//	)
	//stA = append(stA,new(a))
	//st := &SliceType{
	//	SliceInterface: []interface{}{new(a),new(b),new(c)},
	//}
	//st.LoopAndCount()
	//MakeStruct()

}
