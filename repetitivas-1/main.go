package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {
	resultado := funciones.Factorizar(1)
	if resultado != -1 {
		fmt.Printf("%v", resultado)
	}
}
