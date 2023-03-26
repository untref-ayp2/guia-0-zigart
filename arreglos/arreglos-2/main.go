package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {
	vector1 := []float64{1, 2, 3, 4}
	vector2 := []float64{1, 2, 3, 4}

	suma, producto := funciones.SumaYProducto(vector1, vector2)
	fmt.Printf("Suma: %v \n", suma)
	fmt.Printf("Producto escalar: %v  \n", producto)
}
