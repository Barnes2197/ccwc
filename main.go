package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"unicode/utf8"
)

var countBytes, countLines, countWords, countCharacters bool

func main() {

	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a simple CLI application",
		Long:  `MyApp is a longer description of your CLI application.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// This is what happens when no subcommands are specified
			file, err := os.Open(args[0])

			if err != nil {
				panic(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)
			var content []string
			for scanner.Scan() {
				// Retrieve the text of the current line
				nextLine := scanner.Text()

				// Process the line (for this example, just print it)
				content = append(content, nextLine)

			}

			lineCount := 0

			for _, line := range content {
				if line != "" {
					lineCount++
				}
			}

			if countLines {
				printOutput(lineCount, args[0])
				return
			}

			if countBytes {
				printOutput(getByteCount(content), args[0])
				return
			}

			if countWords {
				printOutput(getWordCount(content), args[0])
				return
			}

			if countCharacters {
				printOutput(getChracterCount(content), args[0])
				return
			}

			wordCount := getWordCount(content)
			byteCount := getByteCount(content)

			output := fmt.Sprintf("%d %d %d %s", lineCount, wordCount, byteCount, args[0])

			fmt.Println(output)

		},
	}

	rootCmd.Flags().BoolVarP(&countBytes, "count-bytes", "c", false, "Output the byte count of a text file")
	rootCmd.Flags().BoolVarP(&countLines, "count-lines", "l", false, "Output the line count of a text file")
	rootCmd.Flags().BoolVarP(&countWords, "count-words", "w", false, "Output the word count of a text file")
	rootCmd.Flags().BoolVarP(&countCharacters, "count-characters", "m", false, "Output the character count of a text file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printOutput(count int, filepath string) {
	fmt.Println(fmt.Sprintf("%d %s", count, filepath))
}

func getByteCount(content []string) int {
	byteCount := 0
	for _, line := range content {
		if line != "" {
			byteCount += len(line) + 1
		}
	}
	return byteCount
}

func getWordCount(content []string) int {
	wordCount := 0
	for _, line := range content {
		words := strings.Split(line, " ")
		sum := 0
		for _, word := range words {
			if word != "" {
				sum += 1
			}
		}
		wordCount += sum
	}
	return wordCount
}

func getChracterCount(content []string) int {
	charCount := -1
	for _, line := range content {
		charCount += utf8.RuneCountInString(line) + 1
	}
	return charCount
}
