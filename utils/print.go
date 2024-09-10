package utils

import "fmt"

type User struct {
	Name    string
	Age     uint
	Country string
}

func Print() {
	person := User{
		Name:    "dilan",
		Age:     22,
		Country: "colombia",
	}
	fmt.Println(person.Name)
}
