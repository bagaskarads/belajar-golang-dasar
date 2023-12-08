package main

import "fmt"

type hasName interface {
	getName() string
}

func sayHello(value hasName)  {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string  {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string  {
	return animal.Name
}

func main()  {
	person := Person {Name: "RM"}
	sayHello(person)

	animal := Animal {Name: "Naga"}
	sayHello(animal)
}