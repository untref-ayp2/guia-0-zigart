package funciones

import (
	"fmt"
)

func Polinomio(independiente, coeficiente, coefPrincipal float32) (resultado string) {

	var independentIsNotZero bool = independiente != 0.0
	var coefficientIsNotZero bool = coeficiente != 0.0
	var leadingIsNotZero bool = coefPrincipal != 0.0

	//FIXME refactorizar
	if coefficientIsNotZero && leadingIsNotZero && independentIsNotZero {
		fmt.Printf(" %f + %f X + %f X**2 \n", independiente, coeficiente, coefPrincipal)
	} else if independentIsNotZero && leadingIsNotZero {
		fmt.Printf("%f + %f X**2 \n", independiente, coefPrincipal)
	} else if independentIsNotZero && coefficientIsNotZero {
		fmt.Printf("%f + %f X \n", independiente, coeficiente)
	} else if coefficientIsNotZero && leadingIsNotZero {
		fmt.Printf("%f X + %f X**2 \n", coeficiente, coefPrincipal)
	} else if independentIsNotZero {
		fmt.Printf("%f \n", independiente)
	} else if coefficientIsNotZero {
		fmt.Printf("%f X ", coeficiente)
	} else if leadingIsNotZero {
		fmt.Printf("%f X**2", coefPrincipal)
	}

	return
}
