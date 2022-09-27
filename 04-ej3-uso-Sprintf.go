/*
Usando el código del ejercicio anterior,
En scope a nivel de paquete, asigna los siguientes valores a las tres variables

a x asignale 42
a y asignale “James Bond”
a z asignale true

En main
a) Usa fmt.Sprintf para imprimir todos los VALORES en un solo string.
ASIGNA el valor retornado de TIPO string usando el operador de declaración corta a  la VARIABLE con el IDENTIFICADOR “s”
b) Imprime el valor almacenado por la variable “s”

*/

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%s is agent with %d years. \t active: %v ", y, x, z)
	fmt.Println(s)
}

// Out: James Bond is agent with 42 years. 	 active: true 
