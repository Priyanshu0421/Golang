package main

import "fmt"

func main() {
	hitesh := User{"Hitesh", "hitesh.go.dev", true, 16}

	//fmt.Println("Struct hitesh is: ", hitesh)

	fmt.Printf("details for the hitesh is: %+v\n", hitesh) //  {Name:Hitesh Email:hitesh.go.dev Status:true Age:16}

	fmt.Printf("name is %v , email is %v", hitesh.Name, hitesh.Email)

}

type User struct { // user with capital U to make it publuc struct as we gonna import values from it
	Name   string
	Email  string
	Status bool
	Age    int
}
