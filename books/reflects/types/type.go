package types

import (
	"fmt"
)

type Foo struct {
	Name string
	Age  int
}

func (this Foo) GetName() string {
	return this.Name
}

func (this Foo) GetDisplayName() string {
	return fmt.Sprintf("Name: %s, Age: %d", this.Name, this.Age)
}

func (this *Foo) SetAge(age int) {
	this.Age = age
}
