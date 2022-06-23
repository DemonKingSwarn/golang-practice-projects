package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	demon := User{"d3m0n", "demon@demonkingswarn.ml", true, 19}
	fmt.Println(demon)
	fmt.Printf("demon details are: %+v\n", demon)
	fmt.Printf("Name is %v and email is %v.\n", demon.Name, demon.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
