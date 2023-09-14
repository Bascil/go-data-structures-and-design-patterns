package main

import (
	"log"
	"os"
)

type User struct {
	Name, Email string
}

type Worker struct {
	users []User
}

func NewWorker(users []User) *Worker{
	return &Worker{users: users}
}

func (w *Worker) Find(email string) *User {
	for i := range w.users {
		user := &w.users[i]

		if user.Email == email {
			return user
		}
	}
	return nil
}

var Database = []User{
	{Email: "basilndonga@gmail.com", Name: "Basil Ndonga"},
	{Email: "ndongabasil@gmail.com", Name: "Ndonga Basil"},
}

func main(){
   email := os.Args[1]

   w := NewWorker(Database)
   log.Printf("Looking for %s", email)

   user := w.Find(email)

   if user != nil {
     log.Printf("Email %s is owned by %s", email, user.Name)
   } else {
	log.Printf("Email %s was not found", email)
   }
}