package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Name, Email string
}

type Worker struct {
	users []User
	ch chan *User
	name string
}

func NewWorker(users []User, ch chan *User, name string) *Worker{
	return &Worker{users: users, ch:ch, name: name}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]

		if strings.Contains(user.Email,email) {
			log.Println(">>", w.name)
			w.ch <- user
		}
	}

}

var Database = []User{
	{Email: "basilndonga@gmail.com", Name: "Basil Ndonga"},
	{Email: "ndongabasil@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndonga.basil1@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndonga.basil2@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil1@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil2@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil3@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil4@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil5@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil6@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil7@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil8@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil9@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil10@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil11@gmail.com", Name: "Ndonga Basil"},
	{Email: "ndongabasil12@gmail.com", Name: "Ndonga Basil"},
}

func main(){
   email := os.Args[1]
   
   ch := make(chan *User)

   // shard db -pass first half to first worker and second half to the second worker
   go  NewWorker(Database[:9], ch, "#1").Find(email) //0:9
   go  NewWorker(Database[9:], ch, "#2").Find(email) // 9:max

   log.Printf("Looking for %s", email)

   for {
	select {
	case user := <- ch:
	  log.Printf("Email %s is owned by %s", user.Email, user.Name)
 
	case <- time.After(100 * time.Millisecond): //assume no email found happens after 1s
	 return
 
	}
   }
   
}