package main

import "fmt"

/*
Son aquellas que pueden recibir un numero indefinido de argumentos.
La sintaxis usa como último parámetro una variable con el prefijo ... antes de su tipo, que indica que todos los
valores que se envíen serán almacenados como parte de esta variable.
*/
func main() {
	fmt.Println(Sumar(2))
	fmt.Println(Sumar(2, 2))
	fmt.Println(Sumar(5, 4, 3))
}

func Sumar(numeros ...int) int {
	//Recibe N cantidad de enteros y los suma entre si
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
