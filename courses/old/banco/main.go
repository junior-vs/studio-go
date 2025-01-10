package main

import (
	"fmt"

	"github.com/junior-vs/studio-go/banco/cliente"
)

func main() {

	clienteJunior := cliente.Titular{"Junior", "123.123.123-55", "Dev"}
	fmt.Println(clienteJunior)

}
