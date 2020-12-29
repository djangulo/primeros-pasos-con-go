package main

import (
	"fmt"
)

// Pi valor de la constante π.
const Pi = 3.1416

func main() {
	var r1, r2 = 3.0, 4.0

	var a = area(r1)
	var b = area(r2)

	fmt.Printf("el area de un círculo de radio %.2f es %.2f\n", r1, a)
	fmt.Printf("el area de un círculo de radio %.2f es %.2f\n", r2, b)
}

func area(radio float64) float64 {
	return Pi * radio * radio
}
