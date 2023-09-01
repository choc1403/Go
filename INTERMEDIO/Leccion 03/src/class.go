package main

import "fmt"

// Herencia en Go

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s With age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 25
	ftEmployee.id = 1
	ftEmployee.name = "Juan"
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee.Person)

}
