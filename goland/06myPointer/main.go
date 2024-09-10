package main

import "fmt"

func main() {
	fmt.Println("Welcome to the study of pointers")
	var ptr *int
	myNumber := 23

	ptr = &myNumber

	println("The value of actual ptr is: ", ptr) // o/p -> 0xc000104f20 -> so this give us the actual memory address of the variable
	println("The value of *ptr is: ", *ptr)      // o/p -> 23 -> and this gives us the value that the pointer is carrying

	*ptr = *ptr + 2

	fmt.Println("New value of the myNumber due to the operation we did above: ", myNumber)
}
