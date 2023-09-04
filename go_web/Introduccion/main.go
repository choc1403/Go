package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	Metodos HTTP:
	GET, POST, PUT, DELETE
*/

func main() {
	// Creando rutas en go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor de header") // se podra mostrar en curl -i http://localhost:3000

		fmt.Println(r.URL.RawQuery) // Obtener los argumeros de la URL
		fmt.Println(r.URL.Query())  // Obtiene los argumentos de una URL en forma de map

		name := r.URL.Query().Get("name") // Forma de obtener un argumeto desde el map
		if len(name) != 0 {
			fmt.Println(name)
		}

		fmt.Fprintf(w, "Hola Mundo")
		fmt.Println("El metodo es +" + r.Method) //Conociendo el metodo efectuado
		http.Redirect(w, r, "/docs", 301)        // Redireccionamientos
		// 301 - http.StatusMovedPermanet - mejor colocar constantes
	})

	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Estas en el apartado de documentaciones")
		// Conciendo los metodos en las cuales esta accediendo el usuario
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hola mundo desde el metodo GET")
		case "POST":
			fmt.Fprintf(w, "Hola mundo desde el metodo POST")
		case "PUT":
			fmt.Fprintf(w, "Hola mundo desde el metodo PUT")
		case "DELETE":
			fmt.Fprintf(w, "Hola mundo desde el metodo DELETE")
		default:
			http.Error(w, "Metodo no valido", http.StatusBadRequest)
		}

	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r) //paginas no encontradas
	})

	http.HandleFunc("/error2", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Este es un error", http.StatusConflict)
	})

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		// Generando QUERY
		fmt.Println(r.URL)
		values := r.URL.Query()
		values.Del("otro")
		values.Add("name", "Choc")
		values.Add("course", "Go Web")
		values.Add("Job", "Codigo Facilito")

		r.URL.RawQuery = values.Encode()
		fmt.Println(r.URL)
		fmt.Fprintf(w, "Hola Mundo")
	})

	// Para levantar un servidor en go
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
