package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {
	var ingresado int
	fmt.Printf("Ingrese un valor: \n")
	fmt.Scanln(&ingresado)
	resultado := funciones.EsPrimo(ingresado)
	fmt.Printf("%v", resultado)
}
