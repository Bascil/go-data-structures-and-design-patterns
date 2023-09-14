package main

import "fmt"

func main(){
	user := make(map[string]string)
	user["name"] = "Basil"
	user["role"] = "Dev"

	// alternatively
	// user := map[string]string{
	// 	"name": "Basil",
	// 	"role": "Dev",
	// }
	fmt.Println(user["role"])
}