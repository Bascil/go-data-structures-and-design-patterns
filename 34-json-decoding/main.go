package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	// receive json data as array of bytes
    jsonBytes := []byte(`{"name": "Basil", "age": 35}`)

	// this is a more structured approach, gives you fine grained control
	// use map string interface when you need dirty quick access to json data
	var user User 
	err := json.Unmarshal(jsonBytes, &user) // first param is the byte array followd by the reference

	if err != nil {
		panic(err)
	}

	fmt.Printf("Name is %s\n", user.Name)
}