package main

import "fmt"

func main() {
	// LIFO - Last In First Out

	defer fmt.Println("General Kenobi")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello There")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
