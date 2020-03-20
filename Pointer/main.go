package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("%v, %T\n", b, b)

	// use * to read value of pointer
	fmt.Println(*b)
}
