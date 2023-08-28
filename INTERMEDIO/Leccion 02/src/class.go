package main

import "fmt"

// Constructor

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// 2
	e2 := Employee{
		id:       1,
		name:     "Juan",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	// 3
	e3 := new(Employee)
	e3.id = 2
	e3.name = "Carlos"
	fmt.Printf("%v\n", *e3)

	// 4
	e4 := NewEmployee(3, "Eloi", false)
	fmt.Printf("%v\n", *e4)
}
