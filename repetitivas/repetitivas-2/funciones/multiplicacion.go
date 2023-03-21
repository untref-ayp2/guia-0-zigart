package funciones

func Multiplicar(primerFactor, segundoFactor int) (resultado int) {
	var resultadoAcumulado int = 0

	if segundoFactor > 0 {
		for i := segundoFactor; i > 0; i-- {
			resultadoAcumulado += primerFactor
		}
	} else if segundoFactor < 0 {
		for i := segundoFactor; i < 0; i++ {
			resultadoAcumulado -= primerFactor
		}
	}

	return resultadoAcumulado
}
