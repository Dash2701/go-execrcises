package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions in golang")

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proresult, myMessage := proAdder(2, 6, 7, 8)
	fmt.Println("Pro Result is ", proresult)
	fmt.Println("Pro Message is ", myMessage)

}

func adder(valone int, valtwo int) int {

	return valone + valtwo

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, values := range values {
		total += values
	}
	return total, "Hi Pro result"
}

func greeter() {

	fmt.Println("Namastey from golang")
}

// func greeterTwo() {
// 	fmt.Println("Another greeter method")

// }
