# Aprendiendo GO

### Conceptos clave

1. `func main()` Es el punto de entrada y salida del programa.
2. No puedes tener variables sin utilizar en el código.

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
    _, _ := fmt.Println("Hello, world!") // => Error porque := indica que tienes variables nuevas y no es así en este caso.
    _, _ = fmt.Println("Hello, world!") // => Success
    fmt.Println("Hello, world!") // => Success
}
```
