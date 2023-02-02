package main

import "fmt"

/*
El scope o alcance de una variable, se refiere a dónde está variable existe y tiene una relación con el lugar en el que
fue declarada.
Cómo funciona el scope?
- Dentro de {} que definen un bloque
- En ocasiones se utilizan bloques dentro de bloques, el alcance de la variable es a partir del bloque en el
cual esta fue definida.
- Por lo general se utiliza indentación para separar el alcance de cada bloque.
*/

var var1 = "Este es el nivel 1"

func main() {
	var var2 = "Este es el nivel 2"
	{
		var var3 = "Este es el nivel 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)
	//fmt.Println(var3) // ocasiona error por estar en un bloque distinto
}
