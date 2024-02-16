package main

import "fmt"

func main() {
	fmt.Println("Welcome to Topic Structs")
	//! No inheritance in golang;	No super or parent concept

	var usr1 User
	usr1.Name = "Hiren"
	usr1.age = 18
	usr1.email = "hirenpatel123@gmail.com"

	fmt.Println(usr1.Name)
	fmt.Println(usr1.age)
	fmt.Println(usr1.email)
}

type User struct {
	Name   string
	age    int
	email  string
	status bool
}
