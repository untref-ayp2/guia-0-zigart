package funciones

import (
	"fmt"
)

func Polinomio(coeficientes ...float64) {
	duplicadoCoeficientes := make([]string, len(coeficientes))

	for index, coeficiente := range coeficientes {
		if index == 0 {
			duplicadoCoeficientes[index] = fmt.Sprintf("%v ", coeficiente)
		} else if index == 1 {
			duplicadoCoeficientes[index] = fmt.Sprintf("+ %v X ", coeficiente)
		} else {
			duplicadoCoeficientes[index] = fmt.Sprintf("+ %v X**%v ", coeficiente, index)
		}
		fmt.Printf(duplicadoCoeficientes[index])
	}
}
