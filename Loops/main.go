package main

import "fmt"

func main() {
	// fizzbuzz challenge:
	// print 1 through 100
	// for multiple of 3 print fizz
	// multiple of 5 print buzz
	// multiple of both print fizzbuzz

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("fizzbuzz")
			} else {
				fmt.Println("fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
