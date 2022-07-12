package main

import "fmt"

type Person struct {
	name     string
	lastname string
	age      uint16
}

type Employee struct {
	Person
	code     string
	location string
}

func NewEmploye(code string, location string, p Person) *Employee {
	return &Employee{
		code:     code,
		location: location,
		Person:   p,
	}
}

func main() {
	person := Person{name: "jesus", lastname: "ortiz", age: 23}
	employee := NewEmploye("1212", "lomas", person)
	fmt.Println(*employee)
	
	employee2 := Employee{}
	employee2.name = "jesusito";
	employee2.lastname = "ortizito"
	fmt.Println(employee2)
}
