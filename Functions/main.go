package main

import "fmt"

func main() {
	fmt.Println(greeting("Blues"))
}

func greeting(name string) string {
	return "hello " + name
}
