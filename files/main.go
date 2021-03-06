package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This need to go in a file - demonkingswarn.is-a.dev"

	file, err := os.Create("./myGoFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
