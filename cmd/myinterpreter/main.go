package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/tokenizer"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	line := 1
	var fileTokens []tokenizer.Token
	var fileErrors []tokenizer.Error
	for fileScanner.Scan() {
		lineTokens, lineErrors := tokenizer.TokenizeLine(fileScanner.Text(), line)
		fileTokens = append(fileTokens, lineTokens...)
		fileErrors = append(fileErrors, lineErrors...)
		line++
	}
	fileTokens = append(fileTokens, tokenizer.Token{Type: tokenizer.EOF, Value: ""})

	for _, err := range fileErrors {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
	for _, token := range fileTokens {
		fmt.Printf("%s %s null\n", token.Type.String(), token.Value)
	}
	if len(fileErrors) > 0 {
		os.Exit(65)
	} else {
		os.Exit(0)
	}
}
