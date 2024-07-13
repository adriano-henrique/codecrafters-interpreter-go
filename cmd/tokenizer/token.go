package tokenizer

func TokenizeLine(rawFileContents string, line int) ([]Token, []Error) {
	var tokens []Token
	var errors []Error
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
			errors = append(errors, Error{Type: UNEXPECTED_CHARACTER, Value: string(c), Line: line})
		}
	}
	return tokens, errors
}
