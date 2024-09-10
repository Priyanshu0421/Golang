package main

import "fmt"

// foo := 30   // this will show the error as we are using this (:=) operator outside the method

const LoginToken = "3savdyu" // as the first letter of variable LogicNumber is capital so it means it is public variable and we can use it anywhere

func main() {
	var username string = "Hitesh"
	fmt.Println(username)
	fmt.Printf("username is of type:  %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 127 // unit-> unsigned int means 0-255
	fmt.Println(smallVal)
	fmt.Printf("small value is of type: %T \n", smallVal)

	var smallFloat float32 = 255.9837497349 // it gonna print only 255.98375
	fmt.Println(smallFloat)
	var bigFloat float64 = 255.9837497349 // it gonna print only 255.9837497349
	fmt.Println(bigFloat)

	// default value and some aliases
	var anotherVariable int // default valur it will get will be 0
	fmt.Println(anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString) // it will type empty Stirng

	var name = 'a'
	fmt.Println(name)
	fmt.Printf("name is of type: %T \n", name) // Stirng so sometimes it find it itself no need to write

	numberOfUser := 3000.00 // volarus operator -> :=  -> for declaration + assignment , = -> for assignment only.
	//  for ex var foo int  = 30 is similar to foo:= 30
	//  but we can't use this volarus operator outside the methods

	fmt.Println(numberOfUser)
	fmt.Printf("numberOfUser is of type %T \n", numberOfUser)

	fmt.Println(LoginToken) //  here we are using that public variable that w have declared before
	fmt.Printf("LoginToken is of type %T \n", LoginToken)
}
