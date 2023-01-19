package main

import (
	"fmt"
)

func main() {
	// Aprendiendo operadores aritmeticos

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)

	x := 12
	y := 2

	// Suma
	resultado := x + y

	fmt.Println("El resultado de la suma", resultado)

	// Resta
	resultado = x - y
	fmt.Println("El resultado de la resta", resultado)

	// Multiplicacion
	resultado = x * y
	fmt.Println("El resultado de la multiplicacion", resultado)

	// Division
	resultado = x / y
	fmt.Println("El resultado de la division", resultado)

	// Modulo
	resultado = x % y
	fmt.Println("El resultado del modulo", resultado)

}
