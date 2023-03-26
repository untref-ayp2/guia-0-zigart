package funciones

func UnionArreglos(arreglo1 []int, arreglo2 []int) (union []int) {
	nuevoArreglo := append(arreglo1, arreglo2...)
	return nuevoArreglo
}
func InterseccionArreglos(arreglo1 []int, arreglo2 []int) (union []int) {
	var interseccion []int
	for _, valor1 := range arreglo1 {
		for _, valor2 := range arreglo2 {
			if valor1 == valor2 {
				interseccion = append(interseccion, valor1)
			}
		}
	}
	return interseccion
}
