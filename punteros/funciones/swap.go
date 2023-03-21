package funciones

func Swap(px, py *int) {

	antiguoPx := *px
	*px = *py
	*py = antiguoPx
}
