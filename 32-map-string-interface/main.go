package main

// used when parsing json data
import (
	"encoding/json"
	"fmt"
)

func main(){
	myJsonData := make(map[string]interface{})

	// declare an array of bytes
	data := []byte(`{"name": "Basil"}`)

	// pass its pointer to the unmarshall function
	err := json.Unmarshal(data, &myJsonData) 

	if err != nil {
		panic(err) // check for syntax errors
	}

	fmt.Println(myJsonData["name"])
}