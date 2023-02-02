package main

import (
	"fmt"
	"reflect"
)

// El paquete reflect permite determinar el tipo de una variable
func main() {
	var alumno string = "Edgar"
	var edad int = 32
	var peso float64 = 85.5

	fmt.Println(reflect.TypeOf(alumno))
	fmt.Println(reflect.TypeOf(edad))
	fmt.Println(reflect.TypeOf(peso))
}
