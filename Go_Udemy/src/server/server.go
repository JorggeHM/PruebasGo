package main

import (
	"fmt"
	"net/http"
)

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creando el primer server")
}

func main() {
	http.HandleFunc("/", hola)
	fmt.Println("Servidor iniciando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
