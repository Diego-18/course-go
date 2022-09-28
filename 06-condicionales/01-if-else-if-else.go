package main

import "fmt"

func main() {
	var number int

	fmt.Print("Please, enter your number: ")
	fmt.Scanln(&number)

	if number == 0 {
		fmt.Println("This number is par.")
	}else if  number%2 == 0{
		fmt.Println("This number is neutral.")
	}else{
		fmt.Println("TThis number is odd.")
	}
}
