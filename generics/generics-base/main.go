package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//	Se crea una función genérica (T constraints.Ordered) donde constraints.Ordered contiene variedad de tipos (incluyendo int/float64)
//
// mapFunc es otra función, que será llamada para realizar una lógica sobre el valor (multiplicar por 2)
func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}

	return newValues

}

// Desde el main se invoca a la funcion anterior y se comprueba que se puede llamar con floats/ints u otra variedad de tipos.
func main() {
	result := MapValues([]float64{1.1, 2.2, 3.3}, func(n float64) float64 {
		return n * 2
	})
	fmt.Printf("return: %v\n", result)

}
