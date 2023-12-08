package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	//address1 := Address{"Void", "Shurima", "Rift"}
	//address2 := address1
	//
	//address2.City = "Noxus"
	//fmt.Println(address1)
	//fmt.Println(address2)

	var address1 Address = Address{"Void", "Shurima", "Rift"}
	var address2 *Address = &address1

	address2.City = "Noxus"

	fmt.Println(address1)
	fmt.Println(address2)
}
