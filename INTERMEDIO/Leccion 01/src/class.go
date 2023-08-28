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

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = " Juan X"
	fmt.Printf("%v", e)

}
