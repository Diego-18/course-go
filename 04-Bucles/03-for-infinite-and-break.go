/*
imprimir los numeros del 0 al 9 usando un bucle infinito y break
*/
package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}
}
