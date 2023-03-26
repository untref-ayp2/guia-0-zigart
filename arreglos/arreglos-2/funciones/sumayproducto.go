package funciones

func SumaYProducto(vector1 []float64, vector2 []float64) (suma []float64, producto float64) {

	suma = make([]float64, len(vector1))
	for i := 0; i < len(vector1); i++ {
		suma[i] = vector1[i] + vector2[i]
		producto += vector1[i] * vector2[i]
	}

	return suma, producto
}
