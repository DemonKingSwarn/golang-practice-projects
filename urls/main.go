package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghj456ghb"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myUrl)

	//parsing
	res, _ := url.Parse(myUrl)

	//fmt.Println(res.Scheme)
	//fmt.Println(res.Host)
	//fmt.Println(res.Path)
	//fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qParams := res.Query()
	fmt.Printf("Type of query params are: %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, val := range qParams {
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
