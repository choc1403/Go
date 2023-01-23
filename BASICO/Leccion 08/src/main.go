package main

import "fmt"

func main() {
	//Aprendiendo a utilizar arrays y slice

	//Array
	fmt.Printf("\nArray\n\n")
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	//Slice
	fmt.Printf("\nSlice\n\n")

	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Recorrido de sllices con rage
	fmt.Printf("\n\n\n\n\n\n")

	// Primera forma

	for i, valor := range slice {
		fmt.Print(i, ")", valor, " ")
	}
	fmt.Println("")

	//Segunda forma
	for _, valor := range slice {
		fmt.Print(valor)
	}
	fmt.Println("")

	//Tercera forma
	for i := range slice {
		fmt.Print(i)
	}
	fmt.Println("")

}
