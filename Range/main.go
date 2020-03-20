package main

import "fmt"

func main() {
	numbers := []int{12, 20, 49, 62, 77, 45}

	for i, number := range numbers {
		fmt.Printf("%d - number: %d\n", i, number)
	}

	// not using index
	for _, number := range numbers {
		fmt.Printf("number: %d\n", number)
	}
}
