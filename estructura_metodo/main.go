package main

import "math"

// Los métodos son funciones que se pueden agregar a los tipos personalizados como lo son las estructuras

type Cirulo struct {
	Radio float64
}

func (c Cirulo) Area() float64 {
	// c instancia de la estructura(puede usarse el puntero *Circulo)
	// Area() nombre del método con parámetros en caso de ser necesario
	// float64 tipo de dato del valor que retorna
	return 3.1416 * math.Pow(c.Radio, 2)
}

func main() {
	circulo := Cirulo{Radio: 13}
	println("El área del circulo es: ", circulo.Area())
}
