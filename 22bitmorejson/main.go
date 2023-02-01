package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome ot JSON")
	// EnCodeJson()
	DeCodeJson()
}

func EnCodeJson() {

	lcoCourses := []courses{
		{"Reactjs BootCamp", 299, "lco", "abc123", []string{"web", "dev"}},
		{"Mern BootCamp", 199, "lco", "abc123", []string{"web", "dev", "Full"}},
		{"Python BootCamp", 499, "lco", "abc123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)

}

func DeCodeJson() {

	jsonDataFromWeb := []byte(`{
		"coursename": "Reactjs BootCamp",
		"Price": 299,
		"website": "lco",
		"tags": ["web","dev"]
    }
	`)

	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON  was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}

}
