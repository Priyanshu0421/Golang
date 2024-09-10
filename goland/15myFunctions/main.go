package main

import "fmt"

func main() {
	//greeter   -> // here we're just getting the refrence of the function and we're not executing the function
	greeter() // ->  using this we are actually excuting the function

	greeterTwo()

	//func greeterTwo(){          /// we can't write a function inside a function
	//	fmt.Println("Method Two")
	//}
	//greeterTwo() so it won't work out

	result := adder(3, 5)

	proResult, myMessage := proAdder(2, 4, 5, 6)

	fmt.Println("The proResult is: ", proResult)
	fmt.Println("pro message is: ", myMessage) //

	fmt.Println("The result is: ", result)

}

func greeterTwo() {
	fmt.Println("Method Two on how to Greet a person")
}

func greeter() {
	fmt.Println("namaste from India")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) { // here we wanted to return two values so we did this and put them in the parenthesis
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hello i am adding something" // this is to if u want to return more thna one thing so you can do this
}
