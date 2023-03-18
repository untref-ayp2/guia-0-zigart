package funciones

func Factorizar(numero int) (factorial int) {

	resultado := 1

	for i := 2; i <= numero; i++ {
		resultado *= i
	}

	return resultado
}
