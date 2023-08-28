package main

import "fmt"

/*
El objetivo de una clase es definir una serie de propiedades y metodos
que un objeto puede usar para alcanzar diferentes objetivos

Go utiliza Structs para generar nuevos tipos de datos que se pueden
utilizar para cumplir tareas en especificos

*/

type Employee struct {
	id   int
	name string
}

// Creando un "metodo" Set en Go
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

// Creando un "Metodo" Get en Go
func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	// Viendo la informacion de la variable que acabamos de crear
	fmt.Printf("%v", e)
	fmt.Println("")

	// Asignando Valores
	e.id = 1
	e.name = " Juan X"
	fmt.Printf("%v", e)
	fmt.Println("")

	// Aprobando la utilizacion de metodos en Go
	e.SetId(5)
	e.SetName("Name 2")
	fmt.Printf("%v", e)
	fmt.Println("\n-----------------------")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
