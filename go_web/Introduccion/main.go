package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("Nombre", "Valor del header")
		fmt.Fprintf(w, "Hola Mundo")

		//fmt.Println("El metodo es:", r.Method)
		//http.Redirect(w, r, "/dos", 301)
		/*
			fmt.Println(r.URL.Query())

			// r.URL.RawQuery

			name := r.URL.Query().Get("name")

			if len(name) != 0 {
				fmt.Println(name)
			}

			fmt.Println(r.URL)
			values := r.URL.Query()
			values.Add("name", "eduardo")
			r.URL.RawQuery = values.Encode()
			fmt.Println(r.URL)
		*/

		accessToken := r.Header.Get("access_token")

		if len(accessToken) != 0 {
			fmt.Println(accessToken)
		}

		r.Header.Set("Nombre", "Valor")
	})

	/*
		http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
			//fmt.Fprintf(w, "Holaaaa")
			//http.NotFound(w, r)
			http.Error(w, "Este es un error.", http.StatusConflict)
		})
	*/

	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
