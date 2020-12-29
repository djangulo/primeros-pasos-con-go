# Sintaxis básica

Go es un lenguaje de programación enfocado a creación de sistemas. Es fuertemente tipado, y cuenta con un garbage collector y tiene soporte explicito para programación concurrente.

Los programas de Go se escriben en programas terminados en `.go`. Cada archivo debe de comenzar con una declaración de paquete que indica a qué paquete pertenece dicho archivo.

La declaración de paquetes es seguida de la declaración de importaciones.

Existen 2 tipos de comentarios:

- Comentarios de línea, que comienzan con `//` y se detienen al terminar la línea.
- Comentarios generales, que comienzan con `/*` y terminan con `*/`.

```go
// sintaxis/circulo
package main

import "fmt"


// Pi valor de la constante π.
const Pi = 3.1416

func main() {
    var radio = 3.0

    var area = Pi * radio * radio

    fmt.Printf("el area de un círculo de radio %.2f es %.2f", radio, area)
}
```

La constante `Pi` es una declaracion de paquete, y las variables `radio` y `area` son locales a la función `main`.

La declaración de paquete `Pi` es accesible no solo por el archivo donde se declara, sino también por todos los demás archivos.

Las declaraciones de funciones llevan la palabra clave `func`, un listado de parámetros, y un listado de resultados.

La función `area` encapsula la lógica de cálculo del area de un círculo, permitiendo reusarla.

```go
package main

import "fmt"

const pi = 3.1416

func main() {
	var r1, r2 = 3.0, 4.0

	var a = area(r1)
	var b = area(r2)

	fmt.Printf("el area de un círculo de radio %.2f es %.2f\n", r1, a)
	fmt.Printf("el area de un círculo de radio %.2f es %.2f\n", r2, b)
}

func area(radio float64) float64 {
	return pi * radio * radio
}
```

## Identificadores

En Go, los nombres de las variables, funciones, tipos, constantes y paquetes deben de comenzar con una letra, seguido de letras, guión bajo (`_`), o dígitos. También es sensible a mayúsculas, `Hola` y `hola` son identificadores distintos.

Idiomáticamente, los nombres se escriben en [camelCase](https://es.wikipedia.org/wiki/Camel_case).

Go es compatible con unicode, así que cualquier letra en unicode es válida. El código de abajo es perfectamente válido (disculpan cualquier atrocidad que pude haber cometido con los idiomas):

```go
// sintaxis/saludo/saludo.go
func 今日は(w io.Writer, 名前 string) {
	fmt.Fprintf(w, "今日は, %s", 名前)
}

func доброЈутро(w io.Writer, име string) {
	fmt.Fprintf(w, "добро jутро, %s", име)
}

func buenDía(w io.Writer, nombre string) {
	fmt.Fprintf(w, "Buen día, %s", nombre)
}
```

Go cuenta con 25 palabras clave, las mismas, no pueden ser usadas como identificadores:

```go 
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

Existen además, un grupo de identificadores pre-declarados. Estos nombres no estan reservados


- Tipos:
  
```go
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr
```

- Constantes:
  
```go
    true false iota
```

- Valor cero (en otros lenguajes se conoce como `None`, `NULL`, o `null`):

```go
	nil
```

Funciones:

```go
	append cap close complex copy delete imag len
	make new panic print println real recover
```

## Declaraciones

Una *declaración* une un identificador que no sea `_` con uno de 6 tipos de entidades:

### Constantes

Las constantes se declaran y nunca cambian su valor.

```go
const Pi = 3.1416
const (
    Alice = "alice"
    Bob = "bob"
)
```

### Tipos.

Por el momento, solo veremos sus declaraciones. Pronto veremos que se puede hacer con ellos.

```go
type A int
type b interface{}
type c struct {
    i int32
    I int64
}
```

### Variables

Se usa `var` para declarar una variable de algún tipo. Se puede omitir el tipo, o la asignación, pero no ambos.

```go
var a int
var b int = 23
var c = 23 // tipo int, inferido
var e // incorrecto
```

Go cuenta con un mecanismo de inicialización donde toda variable de tipo simple es inicializada con su valor-cero, esto es

- `false`, para tipo `bool`
- `""`, para `string`
- `0` para todo tipo numérico

Se pueden declarar variables de manera consecutiva:

```go
var a, b, c int // a, b, y c tienen valor 0
var d, e, f int = 1, 2, 3 // d == 1, e == 2, f == 3
```

Incluso se pueden combinar tipos si se omite el tipo en sí

```go
var i, b, c, f = 2, true, 3+2i, 3.1417
```

`i` es tipo `int`, `b` es tipo `bool`, `c` es tipo `complex128`, `f` es tipo `float64`

Se pueden usar declaraciones de variables cortas dentro de una función, con el operador `:=`, y no requieren la palabra clave `var`. El tipo es inferido del valor asignado:

```go
i, b, c, f := 2, true, 3+2i, 3.1417
```

Es importante tener en cuenta que el operador `:=` declara una variable, mientras que el operador `=` es una asignación:

```go
var a int // declaración de variable a, tipo int
a = 3 // asignación a variable a, valor 3
a = 3.0 // incorrecto, no se puede asignar 3.0 (tipo float64) a int
b := 4 // declaracion de variable b, tipo int, y asignación de valor 3 a b
b = 5 // asignación de valor 5 a variable b
var c int = 6 // declaracion+asignación
```

La asignación corta no necesariamente declara variables nuevas. Si una variable ya existe, el operador `:=` reasigna el valor a ellas

En el ejemplo que sigue, `fh` es declarada por la segunda línea, y su valor asignado inmediatamente. `err` es reusada, ya que fue declarada en la primera línea

```go
var err error
fh, err := os.Open("archivo.txt")
```

El guión bajo (`_`) se conoce como identificador en blanco, el mismo se puede usar para declaraciones, pero su valor es ignorado. Es usado generalmente para descartar valores no deseados.

En el ejemplo siguiente, la función `os.Open` retorna 2 valores, un `*os.File`, y un error. En ocasiones, aunque no es buena práctica, se busca ignorar los errores

```go
var fh *os.File
fh, _ = os.Open("archivo.ext")
```

## Punteros

> El valor de un puntero es la  *direccion en memoria* de una variable. Un puntero es la ubicación en donde el valor está guardado. No todo valor tiene una dirección, pero todas las variables tienen una.

— Alan A. A. Donovan, The Go Programming language

Con los punteros es posible accesar y/o modificar los valores de una variable de manera indirecta.

El operador `&T` se lee "la dirección de la variable `T`, y es un puntero a la variable de tipo T.

```go
var a int = 3
b := &a // tipo *int, apunta al valor de a
fmt.Println(b, *b) // 0xc00018c000 1
*b = 2
fmt.Println(a) // 2
```

El valor cero de los punteros es `nil`. Se puede probar si un puntero es nulo con

```go
var b *int
fmt.Println(b == nil) // verdadero
```

Si se le asigna un valor, el puntero no se considera nulo

```go
*b = 3
fmt.Println(b == nil) // falso
```

La función `new` acepta un tipo `T`, y retorna un puntero a dicho tipo, `*T`:

```go
a := new(int32) // a es tipo *int32, valor de a == *a = 0
b := new(bool) // b es tipo *bool, valor de b == *b = falso
```







