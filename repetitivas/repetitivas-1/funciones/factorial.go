package funciones

import (
	"fmt"
)

func Factorizar(numero int) (factorial int) {
	if numero < 0 {
		fmt.Printf("No se puede factorizar un numero menor a cero")
	} else {
		resultado := 1

		for i := 2; i <= numero; i++ {
			resultado *= i
		}

		return resultado
	}
	return -1
}
