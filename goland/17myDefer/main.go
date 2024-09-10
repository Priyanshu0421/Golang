package main

import "fmt"

func main() {
	defer myDefer()
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i) // it will return in the reverse order like -> 4 3 2 1 0
	}
}
