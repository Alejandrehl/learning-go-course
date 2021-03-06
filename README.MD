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
10. No se recomienda utilizar la programación funcional en Go. Los callbacks son funciones que se pasan como parámetro a otro función y son la base de la programación funcional, esto se ocupa muchos en otros lenguajes de programación así que de todas formas sirve estar al tanto. [Explicación](https://www.udemy.com/course/aprende-a-programar-con-go/learn/lecture/12502022#overview).
11. Todo en Go es pasado por valor.
12. Pointers (Punteros): Todos los valores son almacenados en memoria. Cada ubicación en memoria tiene una dirección. Un pointer es una dirección en memoria.
    - `&`: Se accede a la dirección del puntero.
    - `*`: Se accede al valor seteado en esa dirección.
    - `\*int`: Se accede al valor seteado en esa dirección.
    - ¿Cuando utilizar punteros?:
      1. Si tenemos una variable o valor que involucra una cantidad grande de datos y necesitamos pasarla por diferentes partes o funciones de nuestra aplicación lo mejor es utilizar un puntero.
      2. Cuando queremos cambiar directamente un valor que esta almacenado en memoria.
13. JSON: Formato ligero de intercambio de datos.
    - ¿Cuando ocupar Marshall?: Cuando quieres tomar un tipo de dato de Go como por ejemplo una Struct y convertirlo en JSON se debe utilizar `json.Marshal(VALOR)`.
    - ¿Cuando ocupar Unmarshall?: Cuando quieres convertir un JSON en tipo de dato Go se debe utilizar `json.Unmarshal(VALOR)`.

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

- Definción de una función simple: `func foo() { logic code }`.
- Definción de una función compleja: `func (r receptor) identificador(params) retorno(s) { logic code }`
  - [Como usar un receptor](https://www.udemy.com/course/aprende-a-programar-con-go/learn/lecture/12501950#overview)
- Defer (Aplazar): La instrucción `defer` aplaza la ejecución de una función hasta que se retorna la función circundante.
- [Closures](https://www.udemy.com/course/aprende-a-programar-con-go/learn/lecture/12502066#overview)

```go
package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()

	// RESULT:
	// bar
	// foo
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
```

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
2. [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
