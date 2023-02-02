package main

import "fmt"

/*
Existen dos formas de enviar una variable a una función:
- Enviar una copia de su valor.
- Enviar un puntero de esta (su dirección en la memoria).
*/

func main() {
	// por copia
	var altura float32 = 1.70
	fmt.Println("La altura es:", altura, "mts")
	fmt.Println("La altura es:", equivalenciaEnPiesCopia(altura), " pies")
	fmt.Println("La nueva altura es:", altura, "mts")

	// por referencia
	var alturaRef float32 = 1.70

	fmt.Println("La altura es:", alturaRef, "mts")
	fmt.Println("Al envejecer:", conversionEnPiesRef(&alturaRef), "mts")
	fmt.Println("Despues de envejecer:", alturaRef, "mts")
}

func equivalenciaEnPiesCopia(altura float32) float32 {
	// convierte la altura a pies y no modifica el valor original de altura
	altura = altura * 3.28
	return altura
}

func conversionEnPiesRef(altura *float32) float32 {
	// modifica el valor de altura original
	*altura = *altura * 3.28
	return *altura
}
