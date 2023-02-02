package funciones

import (
	"fmt"
	"time"
)

// Para funciones publicas se inicia con Mayúscula su nombre
// Para funciones privadas se inicia con Minúsculas su nombre

func hola(s string) string {
	//función privada que recibe 1 parámetro
	return "Hola " + s
}

func sumar(a int, b int) int {
	// función privada que recibe 2 parámetros y los suma
	return a + b
}

func ImprimirAnio() {
	hoy := time.Now()
	fmt.Printf("current year: %d", hoy.Year())
}
