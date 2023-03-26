package funciones

func Sumador(numeros ...int) (valor int) {
	var total int = 0
	for _, valor := range numeros {

		total += valor
	}
	return total
}
