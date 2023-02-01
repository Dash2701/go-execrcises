package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednessday", "Thursday", "Friday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])

	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("Index is  and value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			goto lco

		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jump")

}
