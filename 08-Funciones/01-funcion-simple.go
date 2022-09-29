package main

import "fmt"

func main() {
	name, age := setPerson()
	greet:= greet(name, age)
	fmt.Println(greet)
}

func greet(name string, age int) string{
	greet:= fmt.Sprintf("Hi %s, my age is: %d. \n", name, age)
	return greet
}

func setPerson()(string, int) {
	var (
		name string
		age int
	)

	fmt.Print("Name: ")
	fmt.Scanln(&name)

	fmt.Print("Age: ")
	fmt.Scanln(&age)

	return name, age
}