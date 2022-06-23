package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on slices")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	//highscores[4] = 777

	highscores = append(highscores, 555, 666, 777, 321)

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
