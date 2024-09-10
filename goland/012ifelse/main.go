package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of the login count -> ")
	loginCount, _ := count.ReadString('\n')

	loginCount = strings.TrimSpace(loginCount)

	i, err := strconv.Atoi(loginCount)

	if err != nil {
		panic(err)
	}

	var result string
	if i < 10 {
		result = "Regular user"
	} else if i > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 loginCount"
	}

	//  will print nothing if else case is not there
	fmt.Println(result)
	num := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of the Number: ")
	newNum, _ := num.ReadString('\n')

	newNum = strings.TrimSpace(newNum) //  To remove the "10\n" and to make it just "10" sp that it can easily be converted into integer
	k, err := strconv.Atoi(newNum)

	if err != nil {
		panic(err) // When a function F calls panic, normal execution of F stops immediately
	} else {
		if k%2 == 0 {
			fmt.Println(k, "is an even number")
		} else {
			fmt.Println(k, "is an odd number")
		}
	}
}
