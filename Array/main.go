package main

import "fmt"

func main() {
	var animalArray [3]string

	animalArray[0] = "monkey"

	fmt.Println(animalArray)

	// declare and assign at the same time
	colorArray := [2]string{"red", "blue"}
	fmt.Println(colorArray)

	// slice
	colorSlice := []string{"brown", "green"}
	fmt.Println(colorSlice)

}
