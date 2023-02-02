package main

import "fmt"

// Los acentos graves (backticks) permiten escribir texto en líneas múltiples
func main() {
	libro := `
    Titulo: Cien Años de Soledad
    Autor: Gabriel García Márquez
    Genero: Novela
    Idioma: Español
    País: Colombia
    Fecha: 1967
    `
	fmt.Println(libro)
}
