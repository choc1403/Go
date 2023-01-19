package main

import "fmt"

func main() {
	// Primer hola mundo desde Go

	fmt.Println("Hola Mundo :)")
	fmt.Println()

	// Declaracion de variables constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("Declaracion de pi1: ", pi, "\nDeclaracion de pi2: ", pi2)
	fmt.Println()

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	area = base * altura

	fmt.Println("La base: ", base, "\nLa altura: ", altura, "\nLa Area: ", area)
	fmt.Println()

	// Tipos de datos
	/*
		int
		string
		float64
		bool
		uint --> Para numeros positivos
		int
		Complex128 --> numeros complejos(Numeros imaginarios)
			c:= 10 + 8i

	*/

}
