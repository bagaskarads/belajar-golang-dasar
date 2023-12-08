package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	var address1 Address = Address{"Void", "Shurima", "Rift"}
	var address2 *Address = &address1
	address2.City = "Noxus"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Zaun", "Freljord", "Rift"}
	fmt.Println(address1)
	fmt.Println(address2)
}