package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("The fruit list is : ", fruitlist)
	fmt.Println("The fruit list is : ", len(fruitlist))

	var veglist = [3]string{"Potato", "beans", "mushrooms"}

	fmt.Println("The veg list is ", veglist)
	fmt.Println("The veg list is ", len(veglist))
}
