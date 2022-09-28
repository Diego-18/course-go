package main

import "fmt"

func main() {

	var fruits [5]string

	// add values to fruits
	fruits[0] = "Apples"
	fruits[1] = "Pears"
	fruits[2] = "Oranges"
	fruits[3] = "Cocos"
	fruits[4] = "Grapes"

	fmt.Println(fruits)

	// array with default values
	vegetables := [3]string{"spinach", "potatoes", "lettuces"}
	fmt.Println(vegetables)
}
