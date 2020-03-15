package main

import (
	"fmt"
	"strconv"
)

var x int = 33

func main() {

	// first way to declare a variable
	var a int
	a = 12

	// second way

	var b int = 62

	// easier way to declare variable
	c := 31

	// innermost scope shadows the package level variable, so x will be 44
	x = 44

	// typecast string into int

	var y string

	y = strconv.Itoa(a)

	fmt.Println(a, b, c, x, y)

}
