package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Hello Welcome to web request modules")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", res)

	defer res.Body.Close() // callers Responsibility to close the connection

	databytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

	rescode := res.StatusCode

	fmt.Println(rescode)

}
