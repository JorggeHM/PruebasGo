package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constantes individuales y en bloque
const pi = 3.14
const (
	x = 100
	y = 0b1010
	z = 0o12
	w = 0xFF
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func papu() {

	fmt.Println()
	fullname := "Jorge Mota \t(alias \"Jorge\")\n"
	fmt.Println(fullname)
}
func main() {
	//Imprimir en c
	// onsola
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())
	papu()

	//Declaracion efectaiva de variables y al final el tipo de dato

	var (
		firstName = "Jorge"
		lastName  = "Mota"
		age       = 23
	)

	var firstName2, lastName2, age2 = "Juan", "Hernandez", 25
	//Imprime variables
	fmt.Println(firstName, lastName, age)
	fmt.Println(firstName2, lastName2, age2)
	fmt.Println(pi, x, y, z, w)
	fmt.Println("Día de la semana:", Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)

}

func init() {
	panic("unimplemented")
}
