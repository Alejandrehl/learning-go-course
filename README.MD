# Aprendiendo GO

### Conceptos clave

1. `func main()` Es el punto de entrada y salida del programa.
2. Go no compila si tienes variables sin utilizar en el código.
3. `:=` Se utiliza para asignar una variable nueva.
4. `=` Se utiliza para asignar un nuevo valor a una variable previamente declarada.
5. `_` Se utiliza para indicar que la función esta retornando algo pero que no se va a utilizar después.
6. `var` se utiliza para definir variables a nível de paquete.
7. Las variables definidas sin la palabra `var` solo pueden ser accedidas dentro de su bloque de código.

```go
var emptyString string // => Returns ""
var name string = "Alejandro"

var defaultZero int // => Returns 0
var age int = 30
```

### Beneficios de usar GO

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
