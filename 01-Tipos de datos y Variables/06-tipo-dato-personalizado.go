/*
Crea tu propio tipo.
Haz que el tipo subyacente, raíz o implícito sea un int.
Crea una VARIABLE de tu nuevo TIPO con el IDENTIFICADOR “x” usando la palabra clave “VAR”

# En func main

a) Imprime el valor de la variable “x”
b) Imprime el tipo de la variable “x”
c) Asigna 42 a la VARIABLE “x” usando el OPERADOR “=”
d) Imprime el valor de la variable “x”
*/
package main

import "fmt"

type cstmEdad int

var x cstmEdad

func main() {
	fmt.Println(x)                                             //a
	fmt.Printf("El tipo de dato de la variable x es: %T\n", x) //b
	x = 42                                                     //c
	fmt.Printf("La edad para obtener beneficios es de: %d", x) //d
}
