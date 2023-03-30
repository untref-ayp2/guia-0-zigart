package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Centro Punto
	Radio  int
}

func (c Circulo) Area() int {
	area := math.Pi * math.Pow(float64(c.Radio), 2)
	return int(area)
}
func (c Circulo) Perimetro() int {
	piPorDos := math.Pi * 2
	radio := float64(c.Radio)
	Perimetro := radio * piPorDos
	return int(Perimetro)
}
func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo: ", c)
	return cadena
}
