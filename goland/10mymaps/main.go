package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the maps")

	languages := make(map[string]string)
	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("Language in the golang is: ", languages) //  map[js:Javascript py:Python rb:Ruby] it will alsi  tell us which datatype it is as we can see here it is showing it is map

	fmt.Println("Js is short for: ", languages["js"]) // Js is short for:  Javascript

	//	 if we want to delete an element in the map

	//delete(languages, "rb")
	//fmt.Println("Language in the golang is: ", languages)   //  map[js:Javascript py:Python]

	//	 Maps in the golang is very intersting thing

	for key, value := range languages {
		fmt.Printf("For key %v -> value is %v\n", key, value)
	}
}
