package funciones

func BuscarMayor(rango []int) (menor, mayor int) {
	var mayorFiltrado int = rango[0]
	var menorFiltrado int = rango[0]
	for index := range rango {
		if mayorFiltrado < rango[index] {
			mayorFiltrado = rango[index]
		}

		if menorFiltrado > rango[index] {
			menorFiltrado = rango[index]
		}
	}
	return menorFiltrado, mayorFiltrado
}
