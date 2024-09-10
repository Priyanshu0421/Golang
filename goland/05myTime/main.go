package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time Study of the Golang")
	//presentTime := time.Now()
	//fmt.Println(presentTime)
	//fmt.Println(presentTime.Format("2006-01-02 15:04:05"))

	createdDate := time.Date(2003, time.September, 29, 12, 0, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
