package main

import "fmt"

var a int

type phone int

var phoneClient phone

func main() {
	a = 45
	fmt.Println(a)		//45
	fmt.Printf("%T\n", a)	//int

	phoneClient = 3015977425
	fmt.Println(phoneClient)	//3015977425
	fmt.Printf("%T\n", phoneClient)	//main.phone

	// Ejemplo de conversion de tipo phone a int
	a = int(phoneClient)
	fmt.Println(a)			//3015977425
	fmt.Printf("%T\n", a)		//int
}
