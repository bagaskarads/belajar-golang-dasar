package main

import "fmt"

func main()  {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Bagas"
	//person["address"] = "Pamulang"

	person := map[string]string {
		"name" : "Bagas",
		"address" : "Pamulang",

	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Bagas"
	book["publisher"] = "Dotoz"

	fmt.Println(book)

	delete(book, "publisher")

	fmt.Println(book)
}