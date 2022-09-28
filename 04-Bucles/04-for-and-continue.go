/*
imprimir los numeros del 0 al 10 usando un bucle y continue
*/
package main

import "fmt"

func main() {
	for index:= 0; index <= 10; index++{
		if index == 5 {
			fmt.Println("We are in the middle of the final value.")
			continue
		}
		fmt.Println(index)
	}
}
