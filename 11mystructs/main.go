package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	amit := User{"Amit", "amit@go.dev", true, 80}
	fmt.Println(amit)
	fmt.Printf("Amit details are %+v\n", amit)
	fmt.Printf("Name is %v and email is %v", amit.Name, amit.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
