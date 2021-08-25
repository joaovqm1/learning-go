package main

import "fmt"

type User struct {
	id int 
	firstName string 
	secondName string 
}

func structs() {
	user1 := User{ id: 1, firstName: "João", secondName: "Vitor" }
	fmt.Println(user1)
}