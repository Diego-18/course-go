package main

import "fmt"

type Course struct {
	nameCourse  string
	urlCourse   string
	skillCourse []string
}

type Career struct {
	nameCareer     string
	durationCareer int
	Course
}

func main() {
	// course1 := Course{
	// 	name:  "Backend fundamentals",
	// 	url:   "http://localhost:8080",
	// 	skill: []string{"Basic", "Advanced", "Expert", "Ninja"},
	// }
	// fmt.Println(course1)

	// course2 := new(Course)
	// course2.name = "Course Go"
	// course2.url = "http://localhost:8080"
	// course2.skill = []string{"Basic", "Advanced", "Expert", "Ninja"}

	Career := new(Career)
	Career.nameCareer = "Software Engineer"
	Career.durationCareer = 5
	Career.nameCourse = "Programming"
	Career.urlCourse = "https://localhost:8080"
	Career.skillCourse = []string{"Back-End", "Front-End", "Data Base"}

	getNamePerson, getLastNamePerson := setPerson()
	// // fmt.Println(course2)
	Career.Registrations(getNamePerson, getLastNamePerson)
}

func setPerson() (string, string) {
	var (
		name     string
		lastName string
	)

	fmt.Print("Name: ")
	fmt.Scanln(&name)

	fmt.Print("LastName: ")
	fmt.Scanln(&lastName)

	return name, lastName
}

func (c Career) Registrations(name string, lastName string) {
	fmt.Printf("The person %s %s has registered in Career for: \t%s\nThe courses you will see are: \t%s\nTo enter you must access the following url: \t%s\nLevels we offer: \t%s\nCareer Duration:\t%d years\n", name, lastName, c.nameCareer, c.nameCourse, c.urlCourse, c.skillCourse, c.durationCareer)
}
