package main

import "fmt"

func Sumar(a int, b int) int {
	return a + b
}

func Restar(a int, b int) int {
	return a - b
}

func Multiplicar(a int, b int) int {
	return a * b
}

func Dividir(a int, b int) int {
	return a / b
}

func main() {
	var a int = 10
	var b int = 5
	fmt.Println("La suma es: ", Sumar(a, b))
}
