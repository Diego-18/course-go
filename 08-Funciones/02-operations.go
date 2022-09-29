package main

import "fmt"

func main() {
	number1, number2 := setNumbers()
	sum := calculateSum(number1, number2)
	sub := calculateSubtract(number1, number2)
	mult := calculateMultiply(number1, number2)
	div := calculateDivision(number1, number2)
	fmt.Println("Result Sum:", sum)
	fmt.Println("Result Sub:", sub)
	fmt.Println("Result Multiply:", mult)
	fmt.Println("Result Division:", div)
}

func setNumbers()(int, int) {
	var (
		number1 int
		number2 int
	)

	fmt.Print("First Number: ")
	fmt.Scanln(&number1)

	fmt.Print("Second Number: ")
	fmt.Scanln(&number2)

	return number1, number2
}

func calculateSum(number1 int, number2 int) int {
	sum := number1 + number2
	return sum
}

func calculateSubtract(number1 int, number2 int) int {
	sub := number1 - number2
	return sub
}

func calculateMultiply(number1 int, number2 int) int {
	mult := number1 * number2
	return mult
}

func calculateDivision(number1 int, number2 int) int {
	div := number1 / number2
	return div
}
