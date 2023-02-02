package main

import "fmt"

type Ciudad struct {
	Nombre    string // Si el nombre está en minúsculas es privada y si es con mayúsculas es publica
	Poblacion int
}

type Pais struct {
	Nombre    string
	Capital   Ciudad
	Idioma    string
	Poblacion int
}

// utilizar estructuras dentro de estructuras
func main() {
	ciudadDeMexico := Ciudad{
		Nombre:    "Ciudad de Mexico",
		Poblacion: 20e6,
	}
	mexico := Pais{
		Nombre:    "México",
		Capital:   ciudadDeMexico,
		Idioma:    "Español",
		Poblacion: 120e6,
	}
	fmt.Printf("%+v\n", mexico)

	// un solo bloque
	mexico2 := Pais{
		Nombre: "México",
		Capital: Ciudad{
			Nombre:    "Ciudad de Mexico",
			Poblacion: 20e6,
		},
		Idioma:    "Español",
		Poblacion: 120e6,
	}
	fmt.Printf("%+v\n", mexico2)
}
