// Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas.

package main

import "fmt"

const (
	Name         string = "Diego"
	lastName     string = "Chavez"
	isProgrammer bool   = true
)

func main() {
	s1 := fmt.Sprintf("My name is %s %s \t Programmer:%t", Name, lastName, isProgrammer)
	fmt.Println(s1)
}
