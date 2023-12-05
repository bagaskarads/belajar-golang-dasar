package main

import "fmt"

func main()  {
	var names [3]string

	names[0] = "RM"
	names[1] = "Bagaskara"
	names[2] = "Aryo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		1, 
		2, 
		3,
		4,
		5,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}