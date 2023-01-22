package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//Conociendo las condicioneles If
	valor1 := 2
	valor2 := 5

	// Solo una operacion
	if valor1 == 3 {
		fmt.Println("El dato del valor 1 es 3")
	} else {
		fmt.Println("El valor real del valor 1 es de ", valor1)
	}

	// Con mas condiciones
	if valor1 == 2 || valor2 == 5 {
		fmt.Println("Exactooooo")
	} else {
		fmt.Println("Nada")
	}

	// Convertir un string en int

	value, err := strconv.Atoi("25")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nValor convertido %v \n", value)

}
