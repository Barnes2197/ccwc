package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a simple CLI application",
		Long:  `MyApp is a longer description of your CLI application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// This is what happens when no subcommands are specified
			readBytes("testfiles/test.txt")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readBytes(filepath string) {
	file, err := os.Open(filepath)

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
