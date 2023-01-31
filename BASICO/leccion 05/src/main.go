package main

import "fmt"

func main() {
	// Estudiando los tipos de ciclos en Go

	// For condicional
	fmt.Printf("For condicional\n")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For while
	fmt.Printf("\n\nFor While\n")
	counter := 0

	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	fmt.Printf("\n\nFor forever\n")
	counterForever := 0

	for {
		counterForever++
		fmt.Println(counterForever)
		if counterForever == 10 {
			break

		}

	}

}
