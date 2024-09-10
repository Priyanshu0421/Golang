package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the loops")
	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	//fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Print(days[d] + " ")
	//}

	//for index := range days {
	//	fmt.Println(days[index])
	//}

	//for index, day := range days {
	//	fmt.Printf("The index value is %v and days is %v\n", index, day)
	//}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco // as soon as rogue value will become 2 it will go to the goto statement we have initialised after the main function
		}

		if rogueValue == 5 {
			rogueValue++ //  will skip the 5 value rest will be the same
			continue
		}

		//if rogueValue == 5 {
		//	break  //  will print value upto 4 only
		//}

		fmt.Println("The Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("This is goto statement and we can call it in our main function as well")
}
