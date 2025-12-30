package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, _ := os.ReadFile("./words.txt")
	_ = data
	fmt.Println("Data in File : " + string(data))
	wordCount := countWords(data)

	fmt.Println("Total Words :", wordCount)

}

func countWords(data []byte) int {
	words := strings.Fields(string(data))
	return len(words)
}
