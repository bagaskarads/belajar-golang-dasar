package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string)  {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main()  {
	var RM Customer
	fmt.Println(RM)

	RM.Name = "RM"
	RM.Address = "GG"
	RM.Age = 22
	fmt.Println(RM)
	fmt.Println(RM.Name)
	fmt.Println(RM.Address)
	fmt.Println(RM.Age)

	BAD := Customer {
		Name: "BAD",
		Address: "WP",
		Age: 22,
	}
	fmt.Println(BAD)

	DS := Customer {"DS", "GN", 22}
	fmt.Println(DS)

	DS.sayHello("Bagas")
	BAD.sayHello("Bagas")
	RM.sayHello("Bagas")
}