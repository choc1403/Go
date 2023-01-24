package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	// Aprendiendo a utilizar los Struct en go

	myCar := car{brand: "Toyota", year: 2017}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}
