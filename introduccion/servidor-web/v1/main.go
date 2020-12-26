// Copyright © 2020 Denis Angulo <djal@tuta.io>
// Licencia: https://creativecommons.org/licenses/by-sa/4.0/deed.es

// servidor-web sirve un saludo en línea.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func hola(w io.Writer, nombre string) {
	if nombre == "" {
		nombre = "mundo"
	}
	fmt.Fprintf(w, "¡Hola %s!\n", nombre)
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	hola(w, q.Get("nombre"))
}
