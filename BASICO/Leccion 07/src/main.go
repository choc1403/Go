package main

import "fmt"

func main() {
	// Aprendiendo a utilizar switch

	modulo := 25 % 2

	switch modulo {
	//este tipo de switch se caracteriza por utilizar como referencia una variable
	case 0:
		fmt.Println("Par")
	default:
		fmt.Println("Impar")
	}

	value := 2
	switch {
	//Este tipo de switch se caracteriza por utilizar multiples casos para diferentes variables
	case value >= 18:
		fmt.Println("Mayor de edad")
	case value <= 0:
		fmt.Println("Dato erronio")
	default:
		fmt.Println("Menor de edad")
	}

}
