package funciones

import "fmt"

func EsPrimo(numero int) (resultado bool) {
	if numero > 1 {
		for i := numero; i > 0; i-- {
			if numero != i && i != 1 && numero%i == 0 {
				return false
			}
		}
		return true
	} else if numero == 1 {
		fmt.Printf("El numero 1 no se considera primo ya que no tiene dos divisores.")
	} else {
		fmt.Printf("El numero debe ser natural.\n")
	}
	return false
}
