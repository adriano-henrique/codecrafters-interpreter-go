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
		case '{':
			tokens = append(tokens, Token{Type: LEFT_BRACE, Value: string(c)})
		case '}':
			tokens = append(tokens, Token{Type: RIGHT_BRACE, Value: string(c)})
		case ',':
			tokens = append(tokens, Token{Type: COMMA, Value: string(c)})
		case '.':
			tokens = append(tokens, Token{Type: DOT, Value: string(c)})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Value: string(c)})
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: string(c)})
		case ';':
			tokens = append(tokens, Token{Type: SEMICOLON, Value: string(c)})
		case '*':
			tokens = append(tokens, Token{Type: STAR, Value: string(c)})
		case '/':
			tokens = append(tokens, Token{Type: SLASH, Value: string(c)})
		default:
			fmt.Printf("Unknown character: %c\n", c)
		}
	}
	tokens = append(tokens, Token{Type: EOF, Value: ""})
	return tokens
}
