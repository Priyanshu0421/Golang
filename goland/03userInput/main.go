package main

import (
	"bufio"
	"os"
)

func main() {
	println("Hello, we will learn how to take user input")

	reader := bufio.NewReader(os.Stdin)
	println("Enter the rating of the pizza")
	// comma ok // err err
	input, _ := reader.ReadString('\n')
	println("Thanks for the rating: ", input)
}
