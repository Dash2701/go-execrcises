package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	logInCount := 10
	var result string

	if logInCount < 10 {
		result = "Regular user"
	} else if logInCount > 10 {
		result = "Watch out"
	} else {
		result = "something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }

}
