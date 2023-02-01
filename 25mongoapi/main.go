package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dash2701/mongoapi/router"
)

func main() {
	fmt.Println("The MongoDB API")
	r := router.Router()
	fmt.Println("Server starting ..")
	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("Listening on port 4001")

}
