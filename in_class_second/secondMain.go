package main

import ("fmt"
"unicode/utf8"
)

var name *string

func main() {
	sum := 1

	for {
		sum += sum
		fmt.Println("Printing out ", sum, " times")
		if sum > 100 {
			break
		}
	}
	fmt.Println("final sum:", sum)
	bytes := len("This 世界")
       fmt.Println("there are ", bytes, "bytes in that string")
	chars := utf8.RuneCountInString("This 世界")	
	fmt.Println(chars)
}
