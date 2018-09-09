package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	jone := person{firstName: "Jone", lastName: "Doe"}

	fmt.Println(jone)
}
