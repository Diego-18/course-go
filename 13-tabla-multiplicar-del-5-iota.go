/* Usando iota, crea constantes para realizar la tabla de multiplicar del 5. Imprime los valores de las constantes.*/
package main

import "fmt"

const (
	_      = 5 * iota
	item1  = 5 * iota
	item2  = 5 * iota
	item3  = 5 * iota
	item4  = 5 * iota
	item5  = 5 * iota
	item6  = 5 * iota
	item7  = 5 * iota
	item8  = 5 * iota
	item9  = 5 * iota
	item10 = 5 * iota
)

func main() {
	fmt.Println(item1)
	fmt.Println(item2)
	fmt.Println(item3)
	fmt.Println(item4)
	fmt.Println(item5)
	fmt.Println(item6)
	fmt.Println(item7)
	fmt.Println(item8)
	fmt.Println(item9)
	fmt.Println(item10)
}
