package main

import (
	"fmt"
	"strconv"
)

// define a struct

type Person struct {
	firstName string
	lastName  string
	age       int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, " + p.firstName + " " + p.lastName + " and my age is " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// init a Person
	person1 := Person{firstName: "Rachel", lastName: "Smith", age: 23}

	// alternative way
	person2 := Person{"Sam", "Rockwell", 44}

	fmt.Println(person1, person2)

	person1.hasBirthday()
	fmt.Println(person1.greet())

}
