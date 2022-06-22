package main

import "fmt"

const LoginToken string = "dghhfjghjdghj" // Public

func main() {
	// fmt.Println("Hello There")
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallFloatVal float64 = 255.4558278236786479
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type: %T \n", smallFloatVal)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "demonkingswarn.is-a.dev"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
