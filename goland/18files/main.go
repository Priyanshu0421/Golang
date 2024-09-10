package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in the golang")

	content := "This need to go in the file -Hello123.com"

	file, err := os.Create("./mygofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("The length is: ", length)

	// THe data you gonna recieve while directly reading the file
	//Readfile("./mygofile.txt")  // [84 104 105 115 32 110 101 101 100 32 116 111 32 103 111 32 105 110 32 116 104 101 32 102 105 108 101 32 45 72 101 108 108 111 49 50 51 46 99 111 109]

	//Here's How we can read the content in the file
	Readfile("./mygofile.txt") //The data in the file is:  This need to go in the file -Hello123.com

	defer file.Close()

}

func Readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // here we rote databye because system read the file in the form of databytes

	checkNilErr(err)

	fmt.Println("The data in the file is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
