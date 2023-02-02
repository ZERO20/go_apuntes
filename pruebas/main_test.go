package main

import "testing"

func TestSumar(t *testing.T) {
	if Sumar(2, 2) != 4 {
		t.Fatal("La suma de 2 + 2 debe ser igual a 4")
	}
}

func TestRestar(t *testing.T) {
	if Restar(10, 2) != 8 {
		t.Fatal("La resta de 10 - 2 debe ser igual a 8")
	}
}

func TestMultiplicar(t *testing.T) {
	if Multiplicar(5, 2) != 10 {
		t.Fatal("La multiplicación de 5 x 2 debe ser igual a 10")
	}
}

func TestDividir(t *testing.T) {
	if Dividir(20, 2) != 10 {
		t.Fatal("La división de 20 / 2 debe ser igual a 10")
	}
}
