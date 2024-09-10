package main

import (
	"fmt"
	"sort"
)

func main() {
	var frulist = []string{"apple", "orange", "banana"}
	fmt.Println(frulist)
	fmt.Printf("The type of the frulist is %T\n", frulist) //  []string

	frulist = append(frulist, "Mango", "peach")
	fmt.Println(frulist) // [apple orange banana Mango peach]

	frulist = append(frulist[1:]) // it will print [orange banana Mango peach]

	frulist = append(frulist[:3])

	fmt.Println(frulist)

	highscores := make([]int, 4)

	highscores[0] = 102
	highscores[1] = 234
	highscores[2] = 867
	highscores[3] = 465
	//highscores[4] = 465  now this will not be added as we make the slices of size 4
	//But if we want to add more value to the slices we can use append method

	highscores = append(highscores, 777, 55, 345)
	fmt.Println(highscores)                     // [102 234 867 465 777 55 345]
	fmt.Println(sort.IntsAreSorted(highscores)) // it will if ints slices are sorted or not  -> false
	sort.Ints(highscores)                       // it will sort the integer slices in the increasing order
	fmt.Println(highscores)                     //[55 102 234 345 465 777 867]
	fmt.Println(sort.IntsAreSorted(highscores)) // true

	// How to remove an element from the slices -> we gonna use the append method
	courses := []string{"reactjs", "javascript", "swift", "ruby", "Nodejs"}

	index := 2

	fmt.Println(courses)

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
