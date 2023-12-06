package main

import "fmt"

func getName() (string, string)  {
	return "RM", "DS"
}

func main()  {
	//firstName, lastName := getName()
	//fmt.Println(firstName, lastName)

	firstName, _ := getName()
	fmt.Println(firstName)
}