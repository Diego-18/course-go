package main

import "fmt"

var (
	name, lastName string
	age, phone     int
)

func main() {
	fmt.Print("Please, enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Please, enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Please, enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Please, enter your phone: ")
	fmt.Scanln(&phone)

	s := fmt.Sprintf("%s %s is programmer with %d years. \t phone: %d ", name, lastName, age, phone)
	fmt.Println(s)
}
