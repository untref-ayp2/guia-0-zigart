package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {
	opciones := []int{1, 2, 7, 15, 20, 3}
	minimo, maximo := funciones.BuscarMayor(opciones)
	fmt.Printf("El valor mínimo es %v \n", minimo)
	fmt.Printf("El valor maximo es %v \n", maximo)
}
