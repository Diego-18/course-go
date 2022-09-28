//Operators

package main

import "fmt"

func main() {
	var (
		number1, number2 int
	)

	// Inputs
	fmt.Print("Enter the first number:")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number:")
	fmt.Scanln(&number2)

	// Process
	sum := number1 + number2
	subtract := number1 - number2
	multiply := number1 * number2
	division := number1 / number2
	mod := number1 % number2

	// Output
	fmt.Println("The sum is:", sum)
	fmt.Println("The subtract is:", subtract)
	fmt.Println("The multiply is:", multiply)
	fmt.Println("The division is:", division)
	fmt.Println("The mod is:", mod)
}
