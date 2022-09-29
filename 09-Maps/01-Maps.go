package main

import "fmt"

func main() {
	listFuits := map[int]string{
		1: "Banana",
		2: "Apple",
		3: "Oranges",
		4: "Pears",
		5: "Cocos",
	}

	listFuits[6] = "Sandia"
	fmt.Println(listFuits)
	fmt.Println(listFuits[4])
	// delete(listFuits, 1)
	// fmt.Println(listFuits)

	listMarvels := make(map[int]string)
	listMarvels[0] = "Captain America"
	listMarvels[1] = "Hulk"
	listMarvels[2] = "Thor"
	listMarvels[3] = "Iron Man"
	listMarvels[4] = "Vision"

	fmt.Println(listMarvels)
}
