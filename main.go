package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []byte
	for scanner.Scan() {
		// Retrieve the text of the current line
		fileBytes := scanner.Bytes()

		// Process the line (for this example, just print it)
		content = append(content, fileBytes...)

	}

	fmt.Println(len(content))

}
