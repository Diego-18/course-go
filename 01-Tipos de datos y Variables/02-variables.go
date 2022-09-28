/*
Usando el operador de declaración corta, ASIGNA los siguientes VALORES a VARIABLES con los IDENTIFICADORES “x”, “y” y “z”

42
“James Bond”
true

Luego imprime los valores almacenados en esas variables usando: 

a) Una sola declaración de la función println
b) Múltiples declaraciones de println

*/

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)		//a

	fmt.Println("El valor de x es: ", x)	//b
	fmt.Println("El valor de y es: ", y)	//b
	fmt.Println("El valor de z es: ", z)	//b
}
