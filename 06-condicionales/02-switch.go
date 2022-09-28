package main

import "fmt"

func main() {
	var (
		number int
		output string
	)

	fmt.Print("Please, enter your number from 1 to 5: ")
	fmt.Scanln(&number)

	switch number {
	case 1:
		output = "first"
	case 2:
		output = "second"
	case 3:
		output = "third"
	case 4:
		output = "fourth"
	case 5:
		output = "fifth"
	default:
		output = "invalid"
	}

	fmt.Println(number, output)
}