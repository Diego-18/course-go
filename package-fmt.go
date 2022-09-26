package main

import "fmt"

var a int
var b string = "test"
var c bool
var d bool = true

func main() {
	e := 42
	f := "un Hola mundo"                          // String Literal Interpretado
	g := `Aprendiendo un nuevo lenguaje llamado GO.` // Row String Literal

	fmt.Printf("El valor de la variable e es: %d\n", e) 		//El valor de la variable e es: 42
	fmt.Printf("El tipo de dato de la variable e es: %T\n", e) //El tipo de dato de la variable e es: int

	fmt.Println(a)                   // 0
	fmt.Println(b)                   // test
	fmt.Println(c)                   // false
	fmt.Println(d)                   // true
	fmt.Println(e)                   // 42
	fmt.Println("Tengo en mi", b, f) // Tengo en mi test un Hola mundo
	fmt.Println(g)                   // Aprendiendo un nuevo lenguaje llamado GO.

	s1 := fmt.Sprint("El ", b, " dice ", f)
	fmt.Println(s1)      //El test dice un Hola mundo
	fmt.Printf("%T", s1) // string

}
