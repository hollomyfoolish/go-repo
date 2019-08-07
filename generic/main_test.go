package generic

import (
	"fmt"
	"testing"
)

type ServiceProvider struct {
	Name            string
	InternalAddress string
	ExternalAddress string
}

type Foo struct{}

func (this Foo) Call() string {
	return "Foo"
}

type FFoo Foo

func (this FFoo) FCall() string {
	return "FFoo"
}

func TestIntArray(t *testing.T) {
	fmt.Println("testing ...")

	arr := IntArray{5, 4, 3, 2, 1}

	fmt.Println(arr)
	Sort(arr)
	fmt.Println(arr)
}

func TestStringArray(t *testing.T) {
	fmt.Println("testing ...")

	arr := StringArray{"abcdef", "abcde", "abcd", "ab", "a"}

	fmt.Println(arr)
	Sort(arr)
	fmt.Println(arr)
}

func TestFoo(t *testing.T) {
	fmt.Println("testing ...")

	foo := FFoo{}

	fmt.Println(foo.FCall())
}
