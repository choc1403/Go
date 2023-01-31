package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	//Goroutines
	var wg sync.WaitGroup

	fmt.Println("Hola")

	wg.Add(1)

	go say("Mundo", &wg)

	wg.Wait()

	go func(texto string) {
		fmt.Println(texto)
	}("Adiooos")

	time.Sleep(time.Second * 1)

}
