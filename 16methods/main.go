package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	amit := User{"Amit", "amit@go.dev", true, 80, 20}
	fmt.Println(amit)
	fmt.Printf("Amit details are %+v\n", amit)
	fmt.Printf("Name is %v and email is %v\n", amit.Name, amit.Email)
	amit.GetStatus()
	amit.NewMail()
	fmt.Printf("Amit details are %+v\n", amit)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "iiitit@rt"
	fmt.Println("Email of this user is :", u.Email)

}
