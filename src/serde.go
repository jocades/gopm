package main

// just testing go and json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		Name: "John Doe",
		Age:  25,
	}

	// Marshal the struct to JSON
	uJSON, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(uJSON))

	// Unmarshal the JSON to a struct
	var u2 User
	err = json.Unmarshal(uJSON, &u2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u2)

	unstrucuted()
}

func unstrucuted() {
	// Unstructured JSON data
	data := []byte(`{"name":"Jane Doe","age":30}`)

	// Unmarshal the JSON to a map
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)

	// convert to user struct from map
	u := User{
		Name: m["name"].(string),
		Age:  int(m["age"].(float64)),
	}

	fmt.Println(u)

}
