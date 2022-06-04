# Aprendiendo GO

Go es un lenguaje creado con una nueva mirada al futuro de la programación. Nuevas palabras para hablar de algunos conceptos. No hablamos de objetos, en Go hablamos de crear tipos y valores de ciertos tipos. No hablamos de casting, hablamos de conversión y aserción.

## Conceptos clave

1. `func main()` Es el punto de entrada y salida del programa.
2. Go no compila si tienes variables sin utilizar en el código.
3. `:=` Se utiliza para asignar una variable nueva.
4. `=` Se utiliza para asignar un nuevo valor a una variable previamente declarada.
5. `_` Se utiliza para indicar que la función esta retornando algo pero que no se va a utilizar después.
6. `var` se utiliza para definir variables a nível de paquete.
7. Las tres palabras usadas para describir buenos nombres de paquetes son: `Corto`, `Conciso` y `Evocativo`.
8. `const` Se utiliza para definir un valor constante, es decir un valor que se setea al momento de la compilación del código.
9. En Go se utilizan `Structs` que son tipos de datos que se componen de otros datos de diferente tipo. Viene siendo lo que en otros lenguajes de programación se conoce como clase.

```go
var emptyString string // => Returns ""
var name string = "Alejandro"

var defaultZero int // => Returns 0
var age int = 30
const fullName string = "Alejandro Exequiel Hernández Lara"
```

7. Las variables definidas sin la palabra `var` solo pueden ser accedidas dentro de su bloque de código.
8. Las funciones que reciben `...interface{}` (Interfaz vacia) significa que reciben cualquier tipo de datos en cualquier cantidad.

## Funciones

- func foo() {}
- func (r receptor) identificador(params) retorno(s) { logic code}

## Beneficios de usar GO

1. El compilador de GO no permite declarar variables no utilizadas, es decir que si declaras una variable y no la estás utilizando el programa va a dar un error de compilación.

```go
package main

import (
	"fmt"
)

func main() {
	bytesNumber, err := fmt.Println("Hello, world!") // => Error por variables no utilizadas.
    _, err := fmt.Println("Hello, world!") // => Error por variable no utilizada.
    _, _ := fmt.Println("Hello, world!") // => Error porque se esta intentando asignar un valor a variables que no se van a utilizar.
    _, _ = fmt.Println("Hello, world!") // => Success
    fmt.Println("Hello, world!") // => Success
}
```

2. En Go se pueden crear nuevos tipos de datos según nuestra necesidad.

```go
package main

import (
	"fmt"
)

var a int

type money int

var b money

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = b
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
```

## Librerías importantes

1. [FMT](https://pkg.go.dev/fmt)
