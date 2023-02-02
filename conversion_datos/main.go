package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// La librería strconv nos permite convertir un tipo de dato a otro
func main() {
	// string a booleano
	var mayorDeEdad string = "true"
	boolVal, _ := strconv.ParseBool(mayorDeEdad)
	fmt.Println(boolVal, reflect.TypeOf(boolVal))

	// booleano a string
	var menorDeEdad bool = true
	strVal := strconv.FormatBool(menorDeEdad)
	fmt.Println(strVal, reflect.TypeOf(strVal))

}
