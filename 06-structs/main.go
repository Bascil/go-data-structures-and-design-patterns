package main

import (
	"errors"
	"fmt"
)


type User struct {
	Name, Role, Email string
	Age int
}

// best practice: when returning multiple arguments, 
// make error the last argument

func (u User) Salary() (int, error) {

	if u.Role == "" {
		return 0, errors.New("Cant handle empty role")
	}

	switch u.Role{
		case "Developer": 
			return 100, nil // no error
		case "Architect": 
		    return 200, nil
	}

	return 0, errors.New(fmt.Sprintf("Unable to locate role '%s'", u.Role))
}

func (u *User) UpdateEmail(email string){
	u.Email = email
}

func main(){
	user := User{Name: "Basil", Role: "Dev", Age: 35}
	//fmt.Println(user.Name,user.Salary())

	// user2 := User{Name: "Josphine", Role: "Architect", Age: 27}
	// fmt.Println(user2.Name,user2.Salary())

	user.UpdateEmail("basilndonga@gmail.com")
	fmt.Println(user)

	// best practice: check errors inline
	if salary, err := user.Salary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}

	// salary, err := user.Salary()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Salary", salary)
	// }
}