package main

import (
	"fmt"
	"log"
)

var name string
var name2 = "Prof Santore"

// this is a comment
func main() {
	//var var3 float64
	//var var4 uint8
	const pi = 3.14159
	name = "CSC386"
	fmt.Println(name, pi)
	fmt.Println(ExampleFunction("sefweFwegfer"))
	length, err := Demo2("wefw	3vw EWCFWEFDwedvfeRF")
	if err != nil {
		log.Fatal("error that shouldn't happen", err)
	}
	fmt.Println("lenght returned by Demo2: ", length)
}

func ExampleFunction(name string) int {
	howLong := len(name)
	if howLong < 5 {
		howLong += 10
	} else {
		howLong = -howLong
	}
	return howLong
}

func Demo2(text string) (int, error) {
	howLong := len(text)
	return howLong, nil
}
