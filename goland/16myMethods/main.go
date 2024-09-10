package main

import "fmt"

func main() {
	hitesh := User{"Hitesh", "hitesh.go.dev", true, 16}

	//fmt.Println("Struct hitesh is: ", hitesh)

	fmt.Printf("details for the hitesh is: %+v\n", hitesh) //  {Name:Hitesh Email:hitesh.go.dev Status:true Age:16}

	fmt.Printf("name is %v , email is %v\n", hitesh.Name, hitesh.Email)

	hitesh.GetStatus()

	hitesh.NewEmail()

	fmt.Printf("name is %v , email is %v\n", hitesh.Name, hitesh.Email)

}

type User struct { // user with capital U to make it publuc struct as we gonna import values from it
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The Status of the user is: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@dev.com"
	fmt.Println("New Email given to the user is: ", u.Email)
}
