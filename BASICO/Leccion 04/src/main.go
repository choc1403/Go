package main

import "fmt"

func funcionNormal() {
	fmt.Println("Holaaaa")
}

func parametros(mensaje string) {
	fmt.Println(mensaje)
}

func multiParametros(a, b int, c string) {
	fmt.Println(a, b, c)
}

func retornoValor(num int) int {
	return num * 2
}

func main() {
	// Uso de funciones

	//normal
	funcionNormal()

	//con parametros
	parametros("Si pasa")

	//Multi parametros
	multiParametros(2, 3, "siiii")

	//retornos de valor
	valor := retornoValor(2)
	fmt.Println(valor)
}
