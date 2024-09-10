package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "orange"
	fruitlist[3] = "peach"

	fmt.Println("The value of the fruitlist is: ", fruitlist)
	fmt.Println("The length of the fruitlist is: ", len(fruitlist))

	var VegetableList = [3]string{"potato", "tomato", "peas"}
	fmt.Println("The value of the VegetableList is: ", VegetableList)
}
