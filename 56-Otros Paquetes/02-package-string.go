package main

import (
	"fmt"
	"strings"
)

func main() {
	var slicen1 []string

	text := "Hi Golang"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	//fmt.Println(strings.Replace(text, "Hi", "Hello", 1))
	fmt.Println(strings.ReplaceAll(text, "Hi", "Hello"))
	fmt.Println(strings.Split(text, " "))

	slicen1 = append(strings.Split(text, " "))
	fmt.Println(slicen1, len(slicen1))
}
