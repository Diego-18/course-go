package main

import "fmt"

func main() {
	course1 := Course{
		name:  "Backend fundamentals",
		url:   "http://localhost:8080",
		skill: []string{"Basic", "Advanced", "Expert", "Ninja"},
	}
	fmt.Println(course1)

	course2 := new(Course)
	course2.name = "Course Go"
	course2.url = "http://localhost:8080"
	course2.skill = []string{"Basic", "Advanced", "Expert", "Ninja"}
	fmt.Println(course2)
}

type Course struct {
	name  string
	url   string
	skill []string
}
