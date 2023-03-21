package main

import (
	"figuras/figuras"
	"figuras/opciones"
	"fmt"
)

func mostrar(f figuras.Figura) {
	fmt.Println(f.ToString())
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}
func mostrarFigura(f figuras.Figura) {
	fmt.Println(f.ToString())
}

func main() {
	p1 := figuras.Punto{0, 0}
	p2 := figuras.Punto{10, 5}

	r := figuras.Rectangulo{P1: p1, P2: p2}
	c := figuras.Cuadrado{Pto: p1, Lado: 10}
	circulo := figuras.Circulo{Centro: p2, Radio: 5}

	arreglo_figuras := [3]figuras.Figura{r, c, circulo}

	for _, f := range arreglo_figuras {
		mostrar(f)
	}

	opciones.MenuDeOpciones()
	fig := opciones.SeleccionarOpciones()
	for _, f := range fig {
		mostrarFigura(f)
	}
}
