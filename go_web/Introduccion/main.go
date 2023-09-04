package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Creando rutas en go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor de header") // se podra mostrar en curl -i localhost:3000
		fmt.Fprintf(w, "Hola Mundo")
		http.Redirect(w, r, "/docs", 301) // Redireccionamientos

	})

	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Estas en el apartado de documentaciones")

	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// Para levantar un servidor en go
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
