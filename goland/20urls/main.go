package main

import (
	"fmt"
	"net/url"
)

var myUrl string = "http://localhost:63343/goland/20urls/newFile.html?_ijt=j51k0mt3ldmg8jak59rg2jbolj&_ij_reload=RELOAD_ON_SAVE"

func main() {

	result, _ := url.Parse(myUrl)
	fmt.Println(result.Host)     // localhost:63343
	fmt.Println(result.Path)     // /goland/20urls/newFile.html
	fmt.Println(result.Scheme)   // http
	fmt.Println(result.RawQuery) // _ijt=j51k0mt3ldmg8jak59rg2jbolj&_ij_reload=RELOAD_ON_SAVE

	qparams := result.Query()

	fmt.Printf("THe type of the query params are: %T\n", qparams) // url.Values

	fmt.Println(qparams["_ij_reload"])

	partsOfUrl := &url.URL{
		Host:     "localhost:63343",
		Path:     "/goland/20urls/newFile.html",
		Scheme:   "http",
		RawQuery: "_ijt=j51k0mt3ldmg8jak59rg2jbolj&_ij_reload=RELOAD_ON_SAVE",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
