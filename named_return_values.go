package main

import "fmt"

func getFullName() (firstName, middleName, lastName string)  {
	firstName = "RM"
	middleName = "Bagaskara"
	lastName = "Aryo"

	return firstName, middleName, lastName
}

func main()  {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}