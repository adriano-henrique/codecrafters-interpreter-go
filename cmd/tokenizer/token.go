package tokenizer

import "fmt"

func Tokenize(rawFileContents string) []Token {
	var tokens []Token
	for _, c := range rawFileContents {
		switch c {
		case '(':
			tokens = append(tokens, Token{Type: LEFT_PAREN, Value: string(c)})
		case ')':
			tokens = append(tokens, Token{Type: RIGHT_PAREN, Value: string(c)})
		default:
			fmt.Printf("Unknown character: %c\n", c)
		}
	}
	tokens = append(tokens, Token{Type: EOF, Value: ""})
	return tokens
}
