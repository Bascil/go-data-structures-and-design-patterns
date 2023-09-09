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
   user := User{"Basil", 35}
   encodedUser, _ := json.Marshal(user)
   fmt.Printf("%s\n",encodedUser)
}