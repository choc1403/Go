package main

import "fmt"

// Interfaces en Go

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
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporatyEmployee) getMessage() string {
	return "Temporaty Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 25
	ftEmployee.id = 1
	ftEmployee.name = "Carlos"
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporatyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)

}
