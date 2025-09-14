package main

import (
	"fmt"
	"net/http"
)

/*func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página principal")
}*/

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de: servidor web en Go")
}

func saludoHandler(w http.ResponseWriter, r *http.Request) {
	nombre := r.URL.Query().Get("nombre")
	if nombre == "" {
		nombre = "anónimo"
	}
	fmt.Fprintf(w, "Hola %s, bienvenido a Go Web!\n", nombre)
}

func main() {
	// Handler básico para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
		//fmt.Fprintln(w, "¡Hola desde Go!")
	})

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/saludo", saludoHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Inicia servidor en el puerto 8080
	fmt.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
