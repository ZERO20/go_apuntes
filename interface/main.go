package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
Una interface es una plantilla de métodos y es utilizado para proporcionar modularidad a Go
- indica que métodos deben ser implementados pero no los define
- son útiles ya que definen la especificación de los métodos, de forma que las posibles implementaciones
pueden ser intercambiadas
*/

type Figura interface {
	Area() float64
	Perimetro() float64
}

// Circulo
type Circulo struct {
	Radio float64
}

// Para implementar la interface Figura implementamos de forma implícita sus métodos Area() y Perímetro
func (c Circulo) Area() float64 {
	return 3.1416 * math.Pow(c.Radio, 2)
}

func (c Circulo) Perimetro() float64 {
	return 2 * 3.1416 * c.Radio
}

// Cuadrado
type Cuadrado struct {
	Lado float64
}

func (c Cuadrado) Area() float64 {
	return math.Pow(c.Lado, 2)
}

func (c Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}
func main() {
	// Creamos instancias para cada estructura y la enviamos como argumento a esta función que imprime los detalles
	figura1 := Circulo{
		Radio: 5.0,
	}
	figura2 := Cuadrado{
		Lado: 8.0,
	}

	ImprimirDetalles(&figura1)
	ImprimirDetalles(&figura2)
}

func ImprimirDetalles(f Figura) {
	/*
		La intención de la modularidad de la interface es que podamos utilizarla como argumento e invocar sus métodos
		independientemente de como estos se implementen. En este caso deseamos imprimir los detalles de
		cada Figura que son Area y Perimetro.
	*/
	fmt.Println("El área es: ", reflect.TypeOf(f).Elem().Name(), f.Area())
	fmt.Println("El perímetro es: ", reflect.TypeOf(f).Elem().Name(), f.Perimetro())
}
