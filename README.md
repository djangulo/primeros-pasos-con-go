# Primeros pasos con `Go`

Aprende a utilizar `Go`, también conocido como Golang.

Este libro cuenta con las siguientes lecciones, a través de las cuales exploraremos diferentes aspectos del lenguaje. Desde la instalación, hasta patrones avanzados.

- [Instalación de Go y configuración de ambiente](https://github.com/djangulo/primeros-pasos-con-go/tree/main/Instalar-go.md)
- [Introducción](https://github.com/djangulo/primeros-pasos-con-go/tree/main/Introduccion.md)
- Sintaxis básica
- Tipos de datos simples
- Tipos de datos compuestos
- Funciones y métodos
- Interfaces
- Organización del código
- Pruebas
- Concurrencia
- Reflexión
- Programación a bajo nivel
- Uso de librerías C con `cgo`

## ¿Cómo usar este libro?

Cada capítulo cuenta con una breve exposición, y uno o más ejemplos ejecutables. Le recomiendo que siga los ejemplos como se exponen, y trate de seguir el código.

Si nunca ha programado antes, la [Introducción](https://github.com/djangulo/primeros-pasos-con-go/Introduccion.md) le puede parecer un tanto intimidante. Le recomiendo que la lea de todas formas, para que vea de lo que Go es capaz.

A partir de [Sintaxis básica](#), las cosas toman un paso mas lento, y no se dejan preguntas sin responder.

## ¿Como está organizado este libro?

Cada capítulo cuenta con su propia carpeta que comparte el mismo nombre. Dentro de dicha carpeta, hay carpetas para cada ejemplo, y la misma puede o no estar separada en versiones, donde cada versión implica una modificación del código.

Por ejemplo, el código del capítulo "Introducción" está organizado como:

```
~$ tree --dirsfirst ch1
introduccion
├── hola-mundo
│   └── hola-mundo.go
├── servidor-web
│   ├── v0
│   │   └── main.go
│   └── v1
│       └── main.go
└── wc
    ├── v0
    │   └── main.go
    └── v1
        └── main.go

7 directories, 5 files
```

## Contribuciones

Someta las correcciones/problemas como un [issue de github](https://github.com/djangulo/primeros-pasos-con-go/issues).

Si quiere someter cambios, o materiales adicionales, cree un [pull request](https://github.com/djangulo/primeros-pasos-con-go/pulls).

Recomendaciones:
- Clone el repositorio, cree un branch, y en dicho branch ponga los materiales que desea incluir.
- El código debera de compilar y correr como su documentación describe. Cree pruebas unitarias para su código.
- El código debe de estar escrito mayormente en Go.
- Trate de atenerse al estilo del resto del libro.
- Sus commits deben de estar [firmados](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/signing-commits).
- Cree su PR temprano, antes de comenzar a escribir código si es posible, con un sujeto explicito. Trate de plantear su plan de acción en la creación del PR. Esto le permitirá a otros proveer retroalimentación sobre el contenido.

## Licencia

<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Licencia Creative Commons" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />Esta obra está bajo una <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Licencia Creative Commons Atribución-CompartirIgual 4.0 Internacional</a>.