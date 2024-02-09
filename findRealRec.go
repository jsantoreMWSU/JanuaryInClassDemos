package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileData, err := os.ReadFile("sillyRec.txt")
	if err != nil {
		log.Fatal("Giving up - can't read file:", err)
	}
	allLines := strings.Split(string(fileData), "\n")
	for lineNum, line := range allLines {
		if lineNum%2 == 0 {
			fmt.Println(line)
		}
		lineNum++
	}

}
