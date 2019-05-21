package gostruct

import (
	"testing"
)

func TestStructToInterface(t *testing.T) {

	type MyStruct struct {
		Field1 string
		Field2 int
		Field3 bool
		Field4 float64
		Field5 float32
	}

	v := MyStruct{
		"Hello",
		100,
		false,
		10.12,
		99.99,
	}

	list, err := ToInterface(v)
	if err != nil {
		t.Fatal(err)
	}

	expectedCount, err := Count(v)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != expectedCount {
		t.Fatal()
	}

}
