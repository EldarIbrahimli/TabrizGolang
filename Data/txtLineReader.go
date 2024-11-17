package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("number of lines:", lineCount)
}
