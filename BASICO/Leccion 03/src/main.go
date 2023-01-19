package main

import "fmt"

func main() {
	// Conociendo el paquete fmt.Print

	hola := "Holaaaa"
	hello := "Que pasoooo"

	fmt.Println(hola, hello)

	nombre := "Plazti"
	cursos := 500

	fmt.Printf("%s Tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v Tiene mas de %v cursos\n", nombre, cursos)

	mensaje := fmt.Sprintf("%s Tiene mas de %d cursos", nombre, cursos)
	fmt.Println(mensaje)

	//Conocer el tipo de dato
	fmt.Printf("Que tipo de dato es la variable nombre %T\n", nombre)

}
