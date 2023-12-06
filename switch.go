package main

import "fmt"

func main()  {
	name := "Bagas"

	switch name {
	case "Bagas":
		fmt.Println("Hello Bagas")
	case "Aryo":
		fmt.Println("Hello Aryo")
	default:
		fmt.Println("WHO???")	
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("xdd")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("xdd")
	case length > 5:
		fmt.Println("xpp")
	default:
		fmt.Println("adhd")
	}
}