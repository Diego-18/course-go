/*
Usa var para DECLARAR tres VARIABLES.
Las variables deben tener scope a nivel de paquete.
No asignar VALORES a las variables.
Usa los siguientes IDENTIFICADORES para las variables y asegúrate de que las variables son de los siguientes TIPOS (lo quiere decir que pueden almacenar VALORES de ese TIPO).

identificador “x” tipo int
identificador “y” tipo string
identificador “z” tipo bool

En main
a) Imprime los valores de cada identificador
b) El compilador asigna valores a las variables. ¿Cómo son llamados esos valores?
*/

package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	//a
	fmt.Println(x) //out: 0
	fmt.Println(y) //out:
	fmt.Println(z) //out: false
}

// b) Los valores son llamados "valor cero"