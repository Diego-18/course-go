package main

import "fmt"

var a int

type phone int

var phoneClient phone

func main() {
	a = 45
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	phoneClient = 3015977425
	fmt.Println(phoneClient)
	fmt.Printf("%T\n", phoneClient)

	// Ejemplo de conversion de tipo phone a int
	a = int(phoneClient)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
