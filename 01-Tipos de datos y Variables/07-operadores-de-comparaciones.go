/*
Usando los siguientes operadores, escribe expresiones y asigna sus valores a variables:
a) ==
b) <=
c) >=
d) !=
e) <
f) >
Imprime los valores de las variables.
*/
package main

import "fmt"

func main() {
	exp1 := (25 == 25)
	exp2 := (12 <= 25)
	exp3 := (12 >= 25)
	exp4 := (52 != 25)
	exp5 := (42 < 20)
	exp6 := (42 > 20)
	fmt.Printf("exp1\t%t\n", exp1) //a
	fmt.Printf("exp2\t%t\n", exp2) //b
	fmt.Printf("exp3\t%t\n", exp3) //c
	fmt.Printf("exp4\t%t\n", exp4) //d
	fmt.Printf("exp5\t%t\n", exp5) //e
	fmt.Printf("exp6\t%t\n", exp6) //f
}
