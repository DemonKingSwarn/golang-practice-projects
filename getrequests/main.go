package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	PerformGetRequest()
}

func PerformGetRequest() {

	const myUrl = "https://demonkingswarn.is-a.dev/anime.html"

	res, err := http.Get(myUrl)

	CheckNilErr(err)

	defer res.Body.Close()

	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length is: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
