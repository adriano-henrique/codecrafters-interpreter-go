package tokenizer

import (
	"bufio"
	"os"
)

func TokenizeFile(readFile *os.File) ([]Token, []Error) {
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	line := 1
	var fileTokens []Token
	var fileErrors []Error
	for fileScanner.Scan() {
		lineTokens, lineErrors := tokenizeLine(fileScanner.Text(), line)
		fileTokens = append(fileTokens, lineTokens...)
		fileErrors = append(fileErrors, lineErrors...)
		line++
	}
	fileTokens = append(fileTokens, Token{Type: EOF, Value: ""})
	return fileTokens, fileErrors
}

func tokenizeLine(rawFileContents string, line int) ([]Token, []Error) {
	var tokens []Token
	var errors []Error
	i := 0
	for i < len(rawFileContents) {
		c := rune(rawFileContents[i])
		peekChar := rune(0)
		if i+1 < len(rawFileContents) {
			peekChar = rune(rawFileContents[i+1])
		}
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
		case '=':
			handlePeekEqual(c, peekChar, &i, &tokens, EQUAL, EQUAL_EQUAL)
		case '!':
			handlePeekEqual(c, peekChar, &i, &tokens, BANG, BANG_EQUAL)
		case '<':
			handlePeekEqual(c, peekChar, &i, &tokens, LESS, LESS_EQUAL)
		case '>':
			handlePeekEqual(c, peekChar, &i, &tokens, GREATER, GREATER_EQUAL)
		case '/':
			if peekChar == '/' {
				return tokens, errors
			} else {
				tokens = append(tokens, Token{Type: SLASH, Value: string(c)})
			}
		case ' ', '\t':
			// Ignore whitespace
		default:
			errors = append(errors, Error{Type: UNEXPECTED_CHARACTER, Value: string(c), Line: line})
		}
		i++
	}
	return tokens, errors
}

func handlePeekEqual(currChar rune, peekChar rune, index *int, tokens *[]Token, singleToken TokenType, doubleToken TokenType) {
	if peekChar == '=' {
		*tokens = append(*tokens, Token{Type: doubleToken, Value: string(currChar) + string(peekChar)})
		*index += 1
	} else {
		*tokens = append(*tokens, Token{Type: singleToken, Value: string(currChar)})
	}
}
