package main

import "fmt"

// Una función recursiva es aquella función que es capaz de invocarse a si misma
func main() {
	ecoDeLaMontana("Yodelayheehoo", 5)
}

func ecoDeLaMontana(mensaje string, iteraciones uint) {
	if iteraciones > 1 {
		ecoDeLaMontana(mensaje, iteraciones-1)
	}
	fmt.Println(mensaje, iteraciones)
}
