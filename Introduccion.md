# Introducción

## ¿Qué es `go` y para qué sirve?

> Go es un lenguaje de programación de fuente abierta, el cual hace fácil construir software simple, confiable y eficiente.
 - del website de Go [golang.org](https://golang.org)

Go se presta especialmente para crear sistemas distribuidos, pero es un lenguaje de uso general, y es usado en dominios como gráficos computacionales, aplicaciones móviles, machine learning y administración de contenedores.

Go puede ser utilizado en una variedad de sistemas y arquitecturas, tanto la familia *nix — Linux, FreeBSD, Mac OS X — como Microsoft Windows, y arquitecturas como i386, amd64, ARM, MIPS. Programas escritos en Go en cualquier combinación de estos ambientes, generalmente funcionan sin ninguna modificación en otro.

A continuación veremos algunos ejemplos ilustrativos utilizando Go. Es fácil ver como estos pudieran extenderse para crear programas más complejos.

Noten que los programas presentados aquí son para una breve exposición de algunas cosas de las cuales Go es capaz. Muchos de los conceptos utilizados aquí serán detallados más adelante.

## ¡Hola Mundo!<a name="hello"></a>

No podemos presentar un lenguaje de programación sin comenzar con el clásico `¡Hola mundo!`.

Cree un archivo llamado `hola-mundo.go`, y escriba el código siguiente.

```go
//  introduccion/hola-mundo
package main

import "fmt"

func main() {
	fmt.Println("¡Hola mundo!")
}
```

y lo corremos con `go run`, desde el directorio donde está el archivo:

```sh
~$ go run hola-mundo.go 
¡Hola mundo!
```

Y si quisieramos compilarlo y convertirlo en un programa, lo podemos hacer con `go build`. Esto crea un archivo binario que puede ser ejecutado sin necesidad de más procesos.

```sh
~$ go build hola-mundo.go
~$ ls
hola-mundo  hola-mundo.go
~$ ./hola-mundo # en windows sería .\hola-mundo.exe
¡Hola mundo!
```

Ahora podemos desglosar el programa, que a pesar de su simplicidad, ilustra muchos puntos clave de el uso de Go.

```go
// introduccion/hola-mundo
package main // 1: declaración de paquete

import "fmt" // 2: importaciones

func main() { // 3: función main
	fmt.Println("¡Hola mundo!")
}
```

1. **Declaración del paquete**: Todo archivo Go debe de 
2. pertenecer a un paquete. El paquete llamado `main` es especial, ya que el mismo le indica al compilador que el codigo que le sigue es para un archivo ejecutable, no para una librería. Esto quiere decir que si el paquete es `main`, debe de haber una función `main` que contiene el código a ejecutar cuando se invoque el programa.
   
   En el [Interfaces](#) veremos como se puede organizar el código en paquetes para dar una estructura lógica y fácil de seguir. 
3. **Importaciones**: Aquí se importan las librerías que se quieren utilizar en nuestro programa. En el caso de `hola-mundo.go`, solo se trata de la librería `fmt`, que contiene utilidades para dar formato a cadenas de caracteres, o strings.
   
   Go cuenta con un sinnúmero de paquetes en su librería estándar, con usos diversos. A lo largo de este manual haremos uso de muchos paquetes de la librería estándar, al igual que algunos mantenidos por la comunidad. También, en ocasiones, crearemos nuestros propios paquetes como abstracciones.

4. **Función main**: La función `main` le indica al compilador que se debe ejecutar en nuestro programa. Puede contener código de uso único, como en el caso de nuestro `hola-mundo.go`, o código de larga vida, como un `socket` o un servidor `http`.

	Las funciones en Go se declaran con la palabra clave `func` seguido del nombre de la función (`main`, en el ejemplo anterior), los parámetros que recibe la función (ninguno, en el caso de `main`), y los tipos de retorno (también vaciós en `main`). En [Funciones y métodos](#) veremos en detalle como declarar funciones, y como dichas funciones impactan nuestro código.

Un punto importante que no  es inmediatamente obvio en el ejemplo, es la cuestión del formato del código. Go le da un enfoque fuerte al formato, y provee una herramienta llamada `gofmt` que reescribe cualquier archivo `.go` en un paquete, quitándole de las manos la responsabilidad a los desarrolladores.

Casi todos los editores de texto se pueden configurar para que que automáticamente corran `gofmt` al guardar el archivo.

Si reescribimos `hola-mundo.go` intencionalmente de manera ilegible, `gofmt` se encarga de dar el formato apropiado:

```go
package main; import "fmt"; func main() {                fmt.Println("¡Hola mundo!"   )}
```

```sh
~ch0/hola-mundo$ gofmt hola-mundo.go
package main

import "fmt"

func main() { fmt.Println("¡Hola mundo!") }
```



## Servidor-web<a name="server"></a>

Uno de los usos más comunes de Go es para servidores y utilidades de redes y comunicación.

La librería `net/http` cuenta con un poderoso servidor que puede ser usado en producción.

```go
// introduccion/servidor-web/v0
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
```

Si corremos este programa, podemos ver que hace exactamente lo que indica: responde un request http con el mensaje "¡Hola mundo!". A continuación lo probamos con `curl` (también puede visitar http://localhost:9090 en su navegador).

```sh
~$ go run main.go & # "&" le dice al shell que corra el programa en el fondo
[1] 1826395 # es el PID del proceso
~$ curl http://localhost:9090
!Hola mundo!
~$ fg # fg lo devuelve al "frente" del shell
go run main.go # ctrl+C para terminarlo
^Csignal: interrupt

```

Pero es un servidor estático, que solo saluda con el mismo string en cada ocasión. Vamos a modificarlo, para volverlo un poco más dinámico.

### Agregando interactividad a nuestro servidor

Escribamos una función, `hola` que acepta un `io.Writer` (no olvide importar `io`) parámetro, `nombre`, la cual usará `fmt.Fprintf` para escribir a dicho `io.Writer`.

Agregaremos una pequeña validación, en caso de que nombre esté vacío (pronto veremos porqué).

```go
// introduccion/servidor-web/v1
...
func hola(w io.Writer, nombre string) {
	if nombre == "" {
		nombre = "mundo"
	}
	fmt.Fprintf(w, "¡Hola %s!\n", nombre)
}

func handler(w http.ResponseWriter, r *http.Request) {
	hola(w, "")
}
```

`io.Writer` es un interfaz, que actúa como un contrato entre los programas. El interfaz `io.Writer` está definido como

```go
// https://pkg.go.dev/io#Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

El mismo, indica que su implementación debe de tener un método llamado `Write`, que acepte un `[]byte`, y retorne el par `(int, error)`.

Los métodos son parecidos a las funciones, excepto que pertenecen a un `struct`. Un struct es un tipo de dato compuesto, el cual cuenta con una serie de atributos. En [Tipos de datos compuestos](#), exploraremos los `structs`, y luego en [funciones y métodos](#), verémos en detalle qué son los métodos.

Si mira la definición de `http.ResponseWriter`, verá que cuenta con el método `Write`:

```go
// https://pkg.go.dev/net/http/#ResponseWriter
type ResponseWriter interface {
	...
	Write([]byte) (int, error)
	...
}
```

Esto implica que el `struct` `ResponseWriter` satisface el interfaz `io.Writer`. En Go, los interfaces se satisfacen implícitamente, esto quiere decir que no es necesario declararlo. Por ejemplo, en `java`, algo similar sería

```java
import io.Writer; // importación del interfaz

public class ResponseWriter implements Writer { // clase que lo implementa
	...
}
```

en Go, basta con escribirle un método al objeto que tenga la misma *firma*. La firma de una función se refiere al nombre de la función, número y tipo de parámetros, y número y tipo de variables de retorno.

Si nunca ha programado anteriormente, muchos de estos términos le parecerán extraños, esto es normal. A lo largo de este libro todos los términos serán explicados.

Si corre el programa, y lo prueba, notará que retorna exactamente el mismo resultado.

Vamos a modificar nuestra función `handler` para lea los valores del `query string`:

```go
// introduccion/servidor-web/v1
...
func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	hola(w, q.Get("nombre"))
}
```

`r.URL.Query()` retorna un `url.Values`, que es un tipo envoltura alrededor de un  `map[string][]string` con los valores del `querystring`. El tipo `map[string][]string` es un arreglo asociativo, donde las llaves son del tipo `string` y los valores son del tipo `[]string`. Los `map`s en Go son similares al `dict` de python, `HashMap` de java, o `Dictionary` de C#.

El método `Get` de `url.Values` retorna el primer valor del `[]string` bajo esa llave, y en caso de no encontrarlo, retorna un string vacío (`""`). Afortunadamente, ya nuestra función `hola` maneja los casos donde el parámetro `nombre` está vacío.

Hasta ahora, el archivo `servidor-web/main.go` se ve así:

```go
// introduccion/servidor-web/v1
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
```

Lo corremos en la línea de comando y lo probamos con `curl` (puede probar en su navegador con http://localhost:9090 y http://localhost:9090/?nombre=Go)

```sh
~$ go run main.go &
[1] 1842582
~$ curl http://localhost:9090
!Hola mundo!
~$ curl http://localhost:9090/?nombre=Go
¡Hola Go!
~$ fg # fg lo devuelve al "frente" del shell
go run main.go # ctrl+C para terminarlo
^Csignal: interrupt
```

Esto es solo una muestra de las capacidades de la familia de paquetes `net/*`. Go es muy comunmente usado para programas en línea, el paquete `net` cuenta con herramientas para trabajar con conexiones via TCP, unix sockets, SMTP, y cookies. Muchas de estas utilidades hacen relativamente trivial la tarea de crear servicios FTP, API REST, interactuar con sistemas de publicación/subscripción (comunmente referidos en ingles como pub/sub).

No es sorpresa que compañías como Google, Netflix, Amazon, Uber y Pinterest, usen Go en algun punto de infraestructura para manejar la interconexión de sus sistemas.


## Línea de comando<a name="cli"></>

Go cuenta con excelentes herramientas para crear programas de línea de comando. Debido a su fácil lectura, velocidad de desarrollo, y versatilidad de uso, en muchas ocasiones es usado para crear herramientas de línea de comandos (e.g. [el CLI de github: https://github.com/cli/cli](https://github.com/cli/cli)).

En el ejemplo siguiente, replicaremos algunas de las utilidades de el programa `wc` ([documentación de `wc`](https://www.gnu.org/software/coreutils/manual/html_node/wc-invocation.html#wc-invocation)), una utilidad de GNU/Linux que te retorna el conteo de líneas, conteo de palabras y conteo de `byte`s en un archivo.

Aquí vemos por primera vez el uso del operador `:=`, el cual, dentro de una función, se puede usar en lugar de la declaración `var` para inicializar una variable, y asignarle un valor. Cuando se usa `:=`, Go hace inferencia del tipo de la variable a partir de su valor

```go
	a := 1.0 // float64
	b := "hola" // string
	c := 1 // int
	var d int
	d = 2
	e = 3 // inválido, la variable e no ha sido declarada
```

A continuación, nuestra primera versión de `wc`:

```go
// introduccion/wc/v0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args es un slice que contiene los argumentos pasados al programa,
	// siendo el primero el programa en sí. Si sólo tiene 1 elemento, se
	// debe a que no nos pasaron el argumento.
	if len(os.Args) < 2 {
		fmt.Println("El argumento es requerido.")
		os.Exit(1)
	}
	// fh File Handle
	fh, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0400)
	if err != nil { // si el error no es nulo, fh lo es
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fh)
	lineas, palabras, bytes := 0, 0, 0
	for scanner.Scan() {
		lineas++
		texto := strings.Trim(scanner.Text(), "\n")
		bytes += len(scanner.Bytes())
		if text != "" {
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
```

Si construimos nuestro programa, y lo corremos sobre el mismo código, vemos el resultado.

La opción `-o` le indica a go build que ruta o nombre darle al archivo binario.

```sh
~$ go build -o wc main.go
~$ ./wc main.go
49 172 1100 main.go
```

Pero `wc` provee más funcionalidad: si no se pasa ningun archivo, o el archivo es `-`, permite leer del `STDIN` como un stream de datos. Vamos a implementarlo.

Declaramos la variable `fh` previamente, y la designamos como el interfaz `io.ReadCloser`, un interfaz que implementa `io.Reader`, con el método `Read(p []byte) (n int, err error)`, y `io.Closer`, con el método `Close() error)`.

Afortunadamente, el objeto que utilizamos, `*os.File`, satisface el interfaz `io.ReadCloser`. Note que `os.Stdin` también es del tipo `*os.File`. El asterisco en `*os.File` indica que dicha variable es un puntero del tipo `os.File`, y que dicha variable (de tipo \*T) contiene la *dirección en memoria* del valor T. Veremos punteros con detalle en [Tipos de datos simples](#), y exploraremos interfaces y satisfacción de los mismos en [Interfaces](#).

Debido a que estamos usando la variable `fh` fuera del alcance de la cláusula `if`, tenemos que declarar la variable `err` con `var err error` *dentro del alcance* del `if`, y al llamar `os.OpenFile(...)` no usamos `:=`, sino `=`. El alcance se refiere al tiempo de vida de las variables, y lo veremos en detalle en [Sintaxis básica](#).

Cambie la parte de la validación por el código siguiente:

```go
// introduccion/wc/v1
...
    var fh io.ReadCloser
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
...
}
```

Y ahora podemos utilizarlo para procesar texto, y PIPE datos del stdin. Cuando esté en modo lectura, presione Ctrl+D para detener el programa.

```sh
~$ go build -o wc main.go
~$ ./wc -
Pablito clavo un clavito                        
Que clavito clavo pablito
2 8 49 -
~$ cat /etc/hosts | ./wc -
6 17 147 -
```

Como hemos visto hasta ahora, la librería estándar de Go ofrece muchas herramientas que facilitan la creación de programas y librerías.

Todavía queda mucho que ver de Go. En los capítulos siguientes exploraremos a fondo los diferentes aspectos del lenguaje.
