package main

import "fmt"

const a = 42
const b = 42.32
const c = "Diego Chavez"
const (
	Youtube = "Diego Chavez Web"
	Github  = "Diego-18"
)

type nombre string

var d nombre

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(c)
	fmt.Printf("%T\n", c)

	d = c
	fmt.Println(d)

	fmt.Println("Mi canal de Youtube es:", Youtube)
	fmt.Println("Mi cuenta de Github es:", Github)
}
