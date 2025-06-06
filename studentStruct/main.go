package main

import "fmt"

//Define a struct for a student
type Student struct{
	Name string
	School string
	Year int
}

func main () {
	//Create a new student
	student1 := Student{
		Name:"Alex",
		School:"FOunders High School",
		Year: 2024,
	}
	//Print student info
	fmt.Println("Student Information:")
	fmt.Println("Name:", student1.Name)
	fmt.Println("School:", student1.School)
	fmt.Println("Year:",student1.Year)

}