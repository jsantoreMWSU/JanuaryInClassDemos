package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputS := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for inputS.Scan() {
		userText := inputS.Text()
		switch userText {
		case "csc386", "CSC386":
			fmt.Println("yes you are here")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("not implemented yet")
		}
		fmt.Print(">")

	}
}
