# Instalación de Go

Las instrucciones oficiales de instalación de Go se encuentran en la [página oficial](https://golang.org/doc/install).

Los métodos más abajo se refieren a los diferentes package manager de cada sistema.

## Mac OS X (>10.9)
    
- Instale `Homebrew`. Desde la terminal:

    ```bash
    ~$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```
    
- Luego instale Go usando `Homebrew`:

    ```bash
    ~$ brew install go
    ```


## Windows

Instale `chocolatey`. Visite https://chocolatey.org/install#individual, y siga las instrucciones. Luego, desde `powershell`:

```powershell
choco install go
```

## Linux

- Utilice su package manager para instalar python:

    - Ubuntu / debian:
    
        ```bash
        ~$ sudo apt get update
        ~$ sudo apt install golang-go
        ```
    
    - Arch:
    
        ```bash
        ~$ sudo pacman -S go
        ```
        
    - Fedora linux:
    
        ```bash
        ~$ sudo dnf install golang
        ```
        
    - RHEL / Centos
    
        ```bash
        ~$ yum module -y install go-toolset
        ```


Una vez instalado, puede verificar que la instalación funciona con:

```bash
~$ go version
go version go1.15.6 linux/amd64
```

# Ambiente de programación de Go

## GOPATH

Por razones históricas, Go configura un solo espacio de trabajo. El mismo puede estar en cualquier parte, pero las instalaciones de Go usan `$HOME/go` por defecto.

La variable de ambiente [`GOPATH`](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable) se utiliza para resolver las declaraciones de importación.

Es importante configurar estas variable, así indicamos de manera explícita donde los queremos. Ponga al final de su `.bash_profile`, `.bashrc` u otro script de inicialización de shell:

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## Módulos de Go

A partir de la versión 1.11 Go cuenta con [`Módulos`](https://github.com/golang/go/wiki/Modules), que pronto será el uso por defecto.

Los módulos permiten, entre otras cosas, tener un proyecto fuera del `GOPATH`, y manejo semi-automatizado de dependencias.

Debajo hay instrucciones para iniciar un proyecto

Para usar módulos, cree una carpeta para su proyecto, y en una terminal o shell

```sh
~$ go mod init nombre-del-proyecto
go: creating new go.mod: module nombre-del-proyecto
```

Para agregar dependencias, sencillamente importelas a su proyecto. Algunas dependencias están versionadas, usted elige la versión al momento de importarlas, sin embargo, no se espera que el nombre del paquete cambie:

```go
...
import "github.com/golang-migrate/migrate" // importa v1
import "github.com/golang-migrate/migrate/v4" // importa v4
```

En el bloque de código anterior, estas importaciones crearían un nombre de conflicto, debido a que ambas versiones exportan el mismo paquete `migrate`.

### Importación de paquetes locales

Los módulos siempre tratan de importar paquetes de repositorios remotos. Si quiere utilizar un paquete local, puede usar la directiva `replace` en su `go.mod`:

Antes:

```diff
module bla

go 1.15

+ replace github.com/djangulo/primeros-pasos-con-go /ruta/a/paquete/local

require github.com/golang-migrate/migrate/v4 v4.14.1 // indirect
```

# Editor de texto<a name="text-editor"></a>

Para escribir los scripts, resulta muy útil tener a mano un editor de texto que provea:

- Intellisense: le completa los comandos.
- Resaltado de sintaxis: ayuda a detectar errores.
- Análisis estático (del código): le indica si hay un problema en su código.

Puede utilizar el que desee, pero recomiendo `Visual Studio Code` <a target="_blank" rel="noopener noreferrer" href="https://code.visualstudio.com/#alt-downloads">https://code.visualstudio.com/#alt-downloads</a>, con el plugin de `Go`: <a target="_blank" rel="noopener noreferrer" href="https://open-vsx.org/extension/golang/Go">https://open-vsx.org/extension/golang/Go</a>.

Una vez instalado, en el menu, seleccione "File">"Open Workspace" y seleccione la carpeta `introduccion-a-python` creada al principio de esta sección.

Otras opciones de editores de texto:
- Visual Studio Code https://code.visualstudio.com/
- Sublime Text https://www.sublimetext.com/
- Xcode https://developer.apple.com/xcode/
- Notepad++ https://notepad-plus-plus.org/

Existen ciertos ambientes integrados de programación (IDE, por sus siglas en inglés) para el desarrollo de Go. Los mismos proveen herramientas como integarciones con `git`, debuggers, herramientas para pruebas, instalación de paquetes, entre otras utilidades. Todo esto, en caso de hacerlo, lo veremos a través de la línea de comandos, presentado de la manera más simple posibel.

El uso de los IDE va más alla de lo que veremos en este libro, y francamente, no lo necesita; al menos no si está iniciando su trayectoria con Go. Aquí un listado en caso de que desee explorarlos:

- Lite IDE: http://liteide.org/en/
- GoLand: https://www.jetbrains.com/go/
- Intellij Idea + Go plugin: https://www.jetbrains.com/idea/
- Eclipse: https://www.eclipse.org/ide/


# Uso básico de línea de comandos<a name="command-line"></a>

A lo largo de este libro haremos uso de la línea de comandos. Aqui hay una lista no exhaustiva de los comandos mas simples y mas usados.

## En OSX / Linux:

| OSX / Linux | Descripción | Ejemplo | Notas |
| :--- |  :--- | :--- | :--- |
| `pwd` | Imprime el directorio actual | `~$ pwd`<br>`/home/djangulo` | -|
| `cd <dir>` |  Cambiar de directorio a `<dir>` | `~$ cd Desktop`<br>`~/Desktop$ pwd`<br>`/home/djangulo/Desktop` | `..` hace referencia al directorio de "arriba": `~$ cd ..` |
| `ls` | Listar archivos en el directorio actual | `~$ ls`<br>`Desktop Documents Downloads` | `ls -la` muestra archivos escondidos y los tabula |
| `man <comando>` | Muestra el manual de el comando que le sigue | -| - |
| `cat ARCHIVO` |  Muestra en pantalla los contenidos de `ARCHIVO`. | - | - |
| `grep <busqueda> <archivo` | Devuelve las líneas de `ARCHIVO` que contienen `<busqueda>` | - | - |
| `rm ARCHIVO` | Borra el archivo `ARCHIVO`. | `~$ rm main.go`  | <span style="color: #aa0000;">PELIGRO: no pasa por bandeja de reciclaje</span> |
| `rm -r DIR` | Borra el directorio `DIR`. | `~$ rm -r wc/` |<span style="color: #aa0000;">PELIGRO: no pasa por bandeja de reciclaje</span> |

## En Windows, con CMD o powershell:

| Windows | Descripción | Ejemplo | Notas |
| :--- | :--- | :--- | :--- |
| `cd` (sin argumentos) | Imprime el directorio actual | `C:\Users\djangulo> cd`<br>`C:\Users\djangulo` | -|
| `cd DIR`| Cambiar de directorio a `DIR` | `C:\Users\djangulo> cd Desktop`<br>`C:\Users\djangulo\Desktop>` | `..` hace referencia al directorio de "arriba": `C:\Users\djangulo> ..` |
| `dir` | Listar archivos en el directorio actual | `C:\Users\djangulo> dir`<br>`12/13/2020 6:31 PM <DIR> Desktop`<br>`12/13/2020 6:31 PM <DIR> Documents`<br>`12/13/2020 6:31 PM <DIR> Downloads` | `ls -la` muestra archivos escondidos y los tabula |
| `type ARCHIVO` | Muestra en pantalla los contenidos de `ARCHIVO`. | - | - |
| `del ARCHIVO` | Borra el archivo `ARCHIVO`. | `C:\Users\djangulo> del main.go`  | <span style="color: #aa0000;">PELIGRO: no pasa por bandeja de reciclaje</span> |
| `rmdir DIR` | Borra el directorio `DIR`. | `C:\Users\djangulo> wc` |<span style="color: #aa0000;">PELIGRO: no pasa por bandeja de reciclaje</span> |

# "Piping" y redireccíon de comandos

Cada programa en la línea de comandos tiene tres flujos de datos que se conectan con el mismo automáticamente:

- `STDIN  (0)` - Entrada estándar (datos alimentados al programa)
- `STDOUT (1)` - Salida estándar (datos impresos por el programa, por defecto hacia la terminal)
- `STDERR (2)` - Error estándar (para mensajes de errores, por defecto hacia la terminal)

Redirección es el medio mediante el cual conectamos estos flujos entre los diferentes programas y comandos.

Notese que la redirección funciona con cualquier programa en la terminal, no solo con los usados en los ejemplos debajo

 - `|` tambien conocido como `pipe`, crea un flujo de datos de un comando a otro. Lo que hace es que pasa los resultados de un comando a el operando del siguiente.
 - `<`, `<<`, `>` y `>>` manejan redirección de/hacia archivos (la flecha apunta en la dirección de los datos). Las flechas dobles (`<<` y `>>`) anexan al archivo, las flechas únicas (`<` y `>`) sobreescriben el archivo completo.
 
 Debajo hay algunos ejemplos:
 
```bash
~$ ls
Desktop Documents Downloads
~$ ls > mis_carpetas.txt
~$ cat mis_carpetas.txt
Desktop Documents Downloads
```

Digamos que en el directorio `Documents` tengo solo 2 archivos, `script1.py` y `script2.py`, y el directorio `Pictures` no existe.
 
```bash
~$ ls Pictures Documents
ls: cannot access 'Pictures': No such file or directory
Documents:
script1.py script2.py
```
 
Se puede enviar el stderr y el stdout a diferentes archivos utilizando sus índices
```bash
~$ ls Pictures Documents 1> stdout.txt 2> stderr.txt
~$ cat stdout.txt
Documents:
script1.py script2.py
~$ cat stderr.txt
ls: cannot access 'Pictures': No such file or directory
```

El "piping" le pasa la salida de un comando a la entrada del otro

```bash
~$ ls
Desktop Documents Downloads
~$ ls | grep Do
Documents
Downloads
```

No hay limite de "piping" se pueden encadenar cuantos comandos se deseen. Aqui un ejemplo que permite buscar que hace la opción de algun comando, en este ejemplo, la opción `-q` del comando `tail`:

- La opcion `-A <numero` de `grep` le dice a `grep` que entrege `<numero>` líneas despues del resultado.
- Se le pone un `backslash` (`\`) antes del guión para "escapar" el caracter, sino `grep` entiende que es una opción para el.

```bash
~$ man tail | cat | grep -A 1 "\-q"
       -q, --quiet, --silent
              never output headers giving file names
```
