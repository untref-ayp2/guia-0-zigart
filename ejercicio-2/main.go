package main

import (
	"fmt"
	"guia/cero/funciones"
)

func main() {
	multipleChoise()
}

func multipleChoise() {
	correcta := funciones.Correcta{Correcta: 5}
	opciones := []int{1, 2, 3, 4, 5, 6, 7}
	if !correcta.ComprobarRangoDelValorCorrecto(opciones) {
		fmt.Printf("~ Debe seleccionar una respuesta correcta que este dentro de sus opciones")
	} else {

		correcta.Opciones(opciones)
		if !correcta.ComprobarRangoDelValorSeleccionado(opciones) {

			fmt.Printf("~ El valor seleccionado no esta en las opciones")

		} else {
			if correcta.ComprobarSiLaRespuestaEsCorrecta() {
				fmt.Printf("Correcto!")
			} else {
				fmt.Printf("Incorrecto!")
			}
		}
	}
}
