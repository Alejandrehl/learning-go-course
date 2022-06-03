# Learning GO Course

### Beneficios de usar GO

1. El compilador de GO no permite declarar variables no utilizadas, es decir que si declaras una variable y no la estás utilizando el programa va a dar un error de compilación.

```go
package main

import (
	"fmt"
)

func main() {
	bytesNumber, err := fmt.Println("Hello, world!")
}
```

El programa escrito anteriormente va a dar un error de compilación ya que bytesNumber y err que son las variables que retorna la función Println no están siendo utilizados.
