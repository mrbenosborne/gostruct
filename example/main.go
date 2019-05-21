package main

import (
	"fmt"

	"github.com/mrbenosborne/gostruct"
)

type myStruct struct {
	Field1 string
	Field2 int
	Field3 bool
	Field4 float64
	Field5 float32
}

func main() {

	v := myStruct{
		"Hello",
		100,
		false,
		10.12,
		99.99,
	}

	list, _ := gostruct.ToInterface(v)
	for _, v := range list {
		fmt.Printf("%v\n", v)
	}

}
