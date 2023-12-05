package main

import "fmt"

func main()  {
	type NPM string

	var myNPM NPM = "111"

	var contoh string = "222"
	var contohNPM NPM = NPM(contoh)

	fmt.Println(myNPM)
	fmt.Println(contohNPM)
}