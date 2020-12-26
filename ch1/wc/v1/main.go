// Copyright © 2020 Denis Angulo <djal@tuta.io>
// Licencia: https://creativecommons.org/licenses/by-sa/4.0/deed.es

// wc retorna el conteo de lineas, palabras, y bytes del archivo.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var fh io.ReadCloser
	// os.Args es un slice que contiene los argumentos pasados al programa,
	// siendo el primero el programa en sí. Si sólo tiene 1 elemento, se
	// debe a que no nos pasaron el argumento.
	if len(os.Args) < 2 || os.Args[1] == "-" {
		fh = os.Stdin
	} else {
		var err error
		fh, err = os.OpenFile(os.Args[1], os.O_RDONLY, 0400)
		if err != nil { // si el error no es nulo, fh lo es
			fmt.Println(err)
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(fh)
	lineas, palabras, bytes := 0, 0, 0
	for scanner.Scan() {
		lineas++
		texto := strings.Trim(scanner.Text(), "\n")
		bytes += len(scanner.Bytes())
		if texto != "" {
			palabras += len(strings.Split(texto, " "))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		fh.Close() // cerramos el archivo
		os.Exit(1)
	}

	fh.Close() // cerramos el archivo
	fmt.Fprintf(os.Stdout, "%d %d %d %s\n", lineas, palabras, bytes, os.Args[1])
	os.Exit(0)
}
