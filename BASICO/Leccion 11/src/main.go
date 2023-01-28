package main

import "fmt"

const (
	Domingo int = iota
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
)

func main() {
	// Aprendiendo a utilizar la secuencia con iota
	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes)

}
