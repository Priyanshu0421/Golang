package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:63343/goland/19webrequest/new.html?_ijt=5irikrmmv7th7r3oee557qbadc&_ij_reload=RELOAD_ON_SAVE"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of the response is: %T\n", response) //  *http.Response ->  we are getting actual memory address of the response not the copy of the response so we can manipulate the response we are getting from the url

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)

	fmt.Println("The content in the url is: ", content)

	defer response.Body.Close()
}
