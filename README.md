# GoStruct
Return an interface from a struct.

[![Build Status](https://travis-ci.org/mrbenosborne/gostruct.svg?branch=master)](https://travis-ci.org/mrbenosborne/gostruct) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/mrbenosborne/gostruct?status.svg)](https://godoc.org/github.com/mrbenosborne/gostruct)

## Installation
```
go get "github.com/mrbenosborne/gostruct"
```

## Example
```go
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

```
