package main

import (
	"fmt"
	"reflect"

	"github.com/hollomyfoolish/go-repo/books/reflects/types"
)

func main() {
	f := types.Foo{
		Name: "Foo",
		Age:  18,
	}

	f.SetAge(20)
	fmt.Printf("%v\n", f.GetDisplayName())

	t := reflect.TypeOf(&f)

	fmt.Printf("%v\n", t)

	fmt.Printf("%v\n", t.Align())
	fmt.Printf("%v\n", t.FieldAlign())
	fmt.Printf("%v\n", t.NumMethod())
	fmt.Printf("%v\n", t.Name())
}
