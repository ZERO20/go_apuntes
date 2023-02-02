package main

import "fmt"

type Pais struct {
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion int
}

/*
Una instancia de una estructura puede ser construida mediante el constructor new.
Esto permite asignar en la memoria el espacio necesario para almacenar un tipo de estructura y
después asígnale los valores a cada uno de sus campos.
*/
func main() {
	mexico := new(Pais)
	fmt.Printf("%+v\n", mexico)
	mexico.Nombre = "México"
	mexico.Capital = "CDMX"
	mexico.Idioma = "Español"
	mexico.Poblacion = 126e6
	fmt.Printf("%+v\n", mexico)
}
