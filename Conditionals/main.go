package main

import "fmt"

func main() {

	// convention is to not use parenthesis in if-else statements
	x := 10
	y := 15

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

}
