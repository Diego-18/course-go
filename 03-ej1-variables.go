/*
Usa var para DECLARAR tres VARIABLES. 
Las variables deben tener scope a nivel de paquete. 
No asignar VALORES a las variables. 
Usa los siguientes IDENTIFICADORES para las variables y asegúrate de que las variables son de los siguientes TIPOS (lo quiere decir que pueden almacenar VALORES de ese TIPO).

identificador “x” tipo int
identificador “y” tipo string
identificador “z” tipo bool

En main imprime los valores de cada identificador
*/

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println("El valor de x es: ", x)
	fmt.Println("El valor de y es: ", y)
	fmt.Println("El valor de z es: ", z)
}
