package file

import "fmt"

type ExampleStruct struct {
	Name string
}

func (x *ExampleStruct) ExampleMethod() {
	fmt.Println("Привет", x.Name)
}
