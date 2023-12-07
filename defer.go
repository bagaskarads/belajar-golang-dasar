package main

import "fmt"

func logging()  {
	fmt.Println("Selesai memanggil funtion")
}

func runApplication()  {
	defer logging()

	fmt.Println("Run application")
}

func main()  {
	runApplication()
}