package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://www.istio.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request to %s: %s", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response code:", resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("The GET request was successful")
	} else {
		fmt.Println("The GET request was not successful")
	}
}
