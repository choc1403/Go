package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.altura * r.base
}

func calcular(f figuras2D) {
	fmt.Println("Area calculada:", f.area())
}

func main() {
	//Aprendiendo a utilizar interfaces y listas de interfaces
	miCuadrado := cuadrado{base: 2}
	miRectangulo := rectangulo{base: 3, altura: 5}

	resultadoC := miCuadrado.area()
	resultadoR := miRectangulo.area()

	fmt.Printf("El area del cuadrado: %v\nEl resultado del Rectangulo: %v\n", resultadoC, resultadoR)

	// Con interfaz
	calcular(miCuadrado)
	calcular(miRectangulo)

	//Lista de interfaces
	miInterfaz := []interface{}{"Hola", 12, true}
	fmt.Println(miInterfaz...)

}
