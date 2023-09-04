package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Creando rutas en go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor de header")
		fmt.Fprintf(w, "Hola Mundo")

	})

	// Para levantar un servidor en go
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
