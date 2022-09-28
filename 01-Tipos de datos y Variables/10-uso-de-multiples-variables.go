package main

import "fmt"

var name, lastName, email string
var age int

func main() {
	name = "Diego"
	lastName = "Chavez"
	email = "ingdiegochavez18@gmail.com"
	age = 25

	fmt.Println("My name is", name, lastName)
	fmt.Println("My email is", email)
	fmt.Println("My age is", age)
}