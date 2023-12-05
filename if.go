package main

import "fmt"

func main()  {
	name := "Aryo"

	if name == "Bagas" {
		fmt.Println("Hallo Bagas")
	} else if name == "Kara" {
		fmt.Println("Hallo Kara")
	} else {
		fmt.Println("Who???????")
	}

	if length := len(name) ; length > 5 {
		fmt.Println("xdd")
	} else {
		fmt.Println("kekw")
	}
}