package main

import "fmt"

func getXdd(name string)string  {
	return "xdd " + name
}

func main()  {
	contoh := getXdd
	misal := getXdd
	fmt.Println(contoh("Bagas"))
	fmt.Println(misal("Bagas"))
}