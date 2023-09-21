package main

import "projectmodule/file"

func main() {
	user1 := &file.ExampleStruct{
		Name: "Stas",
	}

	user1.ExampleMethod()
}
