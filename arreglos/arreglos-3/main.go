package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {
	arreglo1 := []int{1, 2}
	arreglo2 := []int{3, 4}
	resultado := funciones.UnionArreglos(arreglo1, arreglo2)
	fmt.Printf("%v \n", resultado)

	resultado2 := funciones.InterseccionArreglos(arreglo1, arreglo2)
	fmt.Printf("%v \n", resultado2)
}
