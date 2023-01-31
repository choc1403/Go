package main

import "fmt"

// Entrada
func say(texto string, c chan<- string) {
	c <- texto
}

func main() {
	// Channels la forma de como organizar las goroutines
	c := make(chan string, 1)

	fmt.Println("Hola")

	go say("Adios", c)

	fmt.Println(<-c)

}
