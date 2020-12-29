package main

import (
	"fmt"
	"io"
)

func 今日は(w io.Writer, 名前 string) {
	fmt.Fprintf(w, "今日は, %s", 名前)
}

func доброЈутро(w io.Writer, име string) {
	fmt.Fprintf(w, "добро jутро, %s", име)
}

func buenDía(w io.Writer, nombre string) {
	fmt.Fprintf(w, "Buen día, %s", nombre)
}
