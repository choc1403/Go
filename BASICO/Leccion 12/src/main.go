package main

import "fmt"

type pc struct {
	marca  string
	precio int
	ram    int
}

func (myPc pc) ping() {
	fmt.Println(myPc.marca, "WXASAD")

}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

// Metodo src
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo una pc de marca %s con una ram de %d y con un precio de Q%d", myPc.marca, myPc.ram, myPc.precio)
}

func main() {
	// Aprender a utilizar Punteros y Structs
	a := 50
	b := &a //Para conocer la direccion de memoria

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // Con * Conocemos el valor que hay en esa dirrecion de memoria

	*b = 100
	// Con este cambio de dato a la variable *b estamos haciendo referencia a la "a", por lo cual la variable "a"
	//cambia su valor

	fmt.Println(a)

	myPc := pc{marca: "ASUS", precio: 7000, ram: 12}
	fmt.Println(myPc)

	myPc.ping()
	// Duplicamos la ram mediante punteros
	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)

}
