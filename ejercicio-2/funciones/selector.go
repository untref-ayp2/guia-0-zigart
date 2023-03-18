package funciones

import "fmt"

type Correcta struct {
	Correcta     int
	Seleccionado int
}

func (c *Correcta) SetCorrecta(correcta int) {
	c.Correcta = correcta
}

func (c *Correcta) Opciones(opciones []int) {
	fmt.Printf("Seleccione una opción: \n")
	for _, valor := range opciones {
		fmt.Printf("~ opción %v \n", valor)
	}

}

func (c *Correcta) ComprobarSiLaRespuestaEsCorrecta() (correcto bool) {
	return c.Seleccionado == c.Correcta
}

func (c *Correcta) ComprobarRangoDelValorCorrecto(opciones []int) (correcto bool) {
	for index := range opciones {
		if c.Correcta == opciones[index] {
			return true
		}
	}
	return false
}

func (c *Correcta) ComprobarRangoDelValorSeleccionado(opciones []int) (correcto bool) {
	var eleccion int
	fmt.Scanln(&eleccion)
	c.Seleccionado = eleccion
	for index := range opciones {
		if c.Seleccionado == opciones[index] {
			return true
		}
	}
	return false
}
