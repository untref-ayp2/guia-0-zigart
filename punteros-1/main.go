package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {

	var px int = 2
	var py int = 3

	funciones.Swap(&px, &py)

	fmt.Printf("PX: %v, PY: %v", px, py)

}
