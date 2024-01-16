package main

import "fmt"

func main() {
	/* Variables in Go - syntax = KEYWORD var + VARIABLE_NAME + TYPE */
	/* String */
	var userName string = "suvodeep"
	fmt.Println(userName)
	fmt.Printf("%T \n", userName)

	/* Boolean */
	var isLoggedIn bool = true
	fmt.Println("Boolean value is: ", isLoggedIn)
	fmt.Printf("%T \n", isLoggedIn)
}
