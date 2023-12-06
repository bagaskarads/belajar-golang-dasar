package main

import "fmt"

func getAlo(name string)string  {
	alo := "Alo " + name
	return alo
}

func main()  {
	result := getAlo("Bagaskara")
	fmt.Println(result)

	fmt.Println(getAlo("Aryo"))
	fmt.Println(getAlo("DS"))
}