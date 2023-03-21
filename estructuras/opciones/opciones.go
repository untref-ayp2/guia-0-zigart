package opciones

import (
	"figuras/figuras"
	"fmt"
	"strings"
)

func MenuDeOpciones() {
	fmt.Printf("Seleccione 5 opciones: \n A. Rectangulo \n B. Cuadrado \n C. Cirulo \n ")
}

func SeleccionarOpciones() []figuras.Figura {
	var a string
	fmt.Scanln(&a)
	fmt.Printf("DEBE SELECCIONAR 4 OPCIONES MAS \n")
	var b string
	fmt.Scanln(&b)
	fmt.Printf("DEBE SELECCIONAR 3 OPCIONES MAS \n")
	var c string
	fmt.Scanln(&c)
	fmt.Printf("DEBE SELECCIONAR 2 OPCIONES MAS \n")
	var d string
	fmt.Scanln(&d)
	fmt.Printf("DEBE SELECCIONAR 1 OPCIONES MAS \n")
	var e string
	fmt.Scanln(&e)

	seleccionado := [5]string{a, b, c, d, e}
	fig := []figuras.Figura{}
	i := len(seleccionado) - 1
	for i >= 0 {
		Valor := strings.ToUpper(seleccionado[i])
		isA := comparacion(Valor, "A")
		isB := false
		isC := false
		if !isA {
			isB = comparacion(Valor, "B")
			if !isB {
				isC = comparacion(Valor, "C")
			}
		}
		if isA {
			fig = append(fig, CrearRectangulo())
		} else if isB {
			fig = append(fig, CrearCuadrado())
		} else if isC {
			fig = append(fig, CrearCiculo())
		}
		i--
	}
	return fig
}

func comparacion(x, f string) (valor bool) {
	return x == f
}

func CrearCuadrado() (cuadrado figuras.Cuadrado) {
	fmt.Printf("CREANDO CUADRADO \n")
	punto := figuras.Punto{}
	var valor int
	fmt.Printf("Valor del punto X: ")
	fmt.Scanln(&valor)
	punto.X = valor
	fmt.Printf("Valor del punto Y: ")
	fmt.Scanln(&valor)
	punto.Y = valor
	fmt.Printf("Valor del lado: ")
	var lado int
	fmt.Scanln(&lado)
	return figuras.Cuadrado{Pto: punto, Lado: lado}
}
func CrearRectangulo() (rectangulo figuras.Rectangulo) {
	fmt.Printf("CREANDO RECTANGULO \n")
	punto1 := figuras.Punto{}
	var valor int
	fmt.Printf("Valor del 1er punto X: ")
	fmt.Scanln(&valor)
	punto1.X = valor
	fmt.Printf("Valor del 1er punto Y: ")
	fmt.Scanln(&valor)
	punto1.Y = valor
	punto2 := figuras.Punto{}

	var valorEntrada2 int
	fmt.Printf("Valor del 2do punto X: ")
	fmt.Scanln(&valorEntrada2)
	punto2.X = valorEntrada2
	fmt.Printf("Valor del 2do punto Y: ")
	fmt.Scanln(&valorEntrada2)
	punto2.Y = valorEntrada2

	return figuras.Rectangulo{P1: punto1, P2: punto2}
}
func CrearCiculo() (circulo figuras.Circulo) {
	fmt.Printf("CREANDO CIRCULO \n")
	punto := figuras.Punto{}
	var valor int
	fmt.Printf("Valor del centro en X: ")
	fmt.Scanln(&valor)
	punto.X = valor
	fmt.Printf("Valor del centro en Y: ")
	fmt.Scanln(&valor)
	punto.Y = valor
	fmt.Printf("Valor del radio: ")
	var radio int
	fmt.Scanln(&radio)
	return figuras.Circulo{Centro: punto, Radio: radio}
}
