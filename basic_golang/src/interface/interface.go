package main

import "fmt"

type Person struct {
	name string
	age  uint32
}

type Employee struct {
	code, location, checkIn, checkOut string
	salary                            uint32
	*Person
}

func NewPerson(name string, age uint32) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func NewEmploye(code string, location string, checkIn string, checkOut string, salary uint32, p *Person) *Employee {
	return &Employee{
		code:     code,
		location: location,
		checkIn:  checkIn,
		checkOut: checkOut,
		salary:   salary,
		Person:   p,
	}
}

func (e Employee) String() string {
	return fmt.Sprintf("\ncode: %s \nlocation: %s \nchekIn: %s \ncheckOut: %s \nsalary: %d \nPerson: %v",
		e.code, e.location, e.checkIn, e.checkOut, e.salary, *e.Person)

}

func (e Employee) getMessage() string {
	return "Employee"
}

func (p Person) getMessage() string {
	return "Person"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	var p *Person = NewPerson("jesus", 18)
	var e *Employee = NewEmploye("123", "lomas", "12", "12:30", 123, p)
	fmt.Println(*p)
	fmt.Println(*e)
	e.name = "jesus chan"
	getMessage(e)
	getMessage(p)
}
