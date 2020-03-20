package main

import "fmt"

func main() {
	// define a map
	phoneBook := make(map[string]int)
	// alternative
	addressBook := map[string]string{"Alex": "aaa", "Bob": "bbb"}
	fmt.Println(addressBook)

	// assign key-value pairs
	phoneBook["Jon"] = 12
	phoneBook["Elyse"] = 91
	phoneBook["Gone"] = 39

	fmt.Println(phoneBook)

	// delete from map
	delete(phoneBook, "Gone")
	fmt.Println(phoneBook)

}
