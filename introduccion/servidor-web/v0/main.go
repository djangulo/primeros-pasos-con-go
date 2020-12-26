// Copyright © 2020 Denis Angulo <djal@tuta.io>
// Licencia: https://creativecommons.org/licenses/by-sa/4.0/deed.es

// servidor-web sirve un saludo en línea.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "¡Hola mundo!\n")
}
