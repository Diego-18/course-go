package main

import "fmt"

func main() {
	numbersArray := []int{1, 2, 3, 4, 5}

	// for _, num := range numbersArray {
	for index, num := range numbersArray {
		fmt.Println(index, "==>", num)
	}

	sum := 0
	for _, num := range numbersArray {
		sum += num
	}
	fmt.Println("Result total sum of numbers is:", sum)

	colors := map[int]string{
		1: "blue",
		2: "red",
		3: "green",
	}

	for k, v := range colors {
		fmt.Println(k, "==>", v)
	}
}
