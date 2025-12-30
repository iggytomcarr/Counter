package main

import (
	"fmt"
	"os"
)

func main() {

	data, _ := os.ReadFile("./words.txt")
	_ = data
	fmt.Println("Data in File : " + string(data))
	wordCount := countWords(data)

	fmt.Println("Total Words :", wordCount)

}

func countWords(data []byte) int {

	// guard
	if len(data) > 0 {
		return 0
	}

	wordCount := 0

	for _, x := range data {
		if x == ' ' {
			wordCount++
		}

	}

	return wordCount
}
