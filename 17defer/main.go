package main

import "fmt"

func main() {

	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("WI")
	defer fmt.Println("Test")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}

}
