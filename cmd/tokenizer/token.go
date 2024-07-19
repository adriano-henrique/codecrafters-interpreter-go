package tokenizer

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
	fileTokens = append(fileTokens, Token{Type: EOF, StringValue: "", Value: "null"})
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
			tokens = append(tokens, Token{Type: LEFT_PAREN, StringValue: string(c), Value: "null"})
		case ')':
			tokens = append(tokens, Token{Type: RIGHT_PAREN, StringValue: string(c), Value: "null"})
		case '{':
			tokens = append(tokens, Token{Type: LEFT_BRACE, StringValue: string(c), Value: "null"})
		case '}':
			tokens = append(tokens, Token{Type: RIGHT_BRACE, StringValue: string(c), Value: "null"})
		case ',':
			tokens = append(tokens, Token{Type: COMMA, StringValue: string(c), Value: "null"})
		case '.':
			tokens = append(tokens, Token{Type: DOT, StringValue: string(c), Value: "null"})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, StringValue: string(c), Value: "null"})
		case '+':
			tokens = append(tokens, Token{Type: PLUS, StringValue: string(c), Value: "null"})
		case ';':
			tokens = append(tokens, Token{Type: SEMICOLON, StringValue: string(c), Value: "null"})
		case '*':
			tokens = append(tokens, Token{Type: STAR, StringValue: string(c), Value: "null"})
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
				tokens = append(tokens, Token{Type: SLASH, StringValue: string(c), Value: "null"})
			}
		case ' ', '\t':
			// Ignore whitespace
		case '"':
			var stringLiteral string
			j := i + 1
			for j < len(rawFileContents) {
				if rawFileContents[j] != '"' {
					stringLiteral += string(rawFileContents[j])
				} else {
					break
				}
				j++
			}
			if j == len(rawFileContents) && rawFileContents[j-1] != '"' {
				errors = append(errors, Error{Type: UNTERMINATED_STRING, Value: "null", Line: line})
			} else {
				tokens = append(tokens, Token{Type: STRING, StringValue: fmt.Sprintf("\"%s\"", stringLiteral), Value: stringLiteral})
			}
			i = j
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var auxNumberList []string
			var numberLiteral string
			j := i
			for j < len(rawFileContents) {
				if !unicode.IsDigit(rune(rawFileContents[j])) && rawFileContents[j] != '.' {
					auxNumberList = append(auxNumberList, numberLiteral)
					numberLiteral = ""
					break
				}
				if rawFileContents[j] == '.' {
					auxNumberList = append(auxNumberList, numberLiteral)
					numberLiteral = ""
					auxNumberList = append(auxNumberList, string(rawFileContents[j]))
				}
				if unicode.IsDigit(rune(rawFileContents[j])) {
					numberLiteral += string(rawFileContents[j])
				}
				j++
			}
			if numberLiteral != "" {
				auxNumberList = append(auxNumberList, numberLiteral)
			}
			i = j
			fmt.Println(auxNumberList)
		default:
			errors = append(errors, Error{Type: UNEXPECTED_CHARACTER, Value: string(c), Line: line})
		}
		i++
	}
	return tokens, errors
}

func handleNumberTokenization(rawFileContents string, currentIndex int) ([]Token, []Error, int) {
	var endIndex int
	auxNumberList, endIndex := buildNumberTokenList(currentIndex, rawFileContents)
	fmt.Println(auxNumberList)
	fmt.Println(endIndex)
	return []Token{}, []Error{}, endIndex
}

func buildNumberTokenList(currentIndex int, rawFileContents string) ([]string, int) {
	var auxNumberList []string
	var numberLiteral string
	j := currentIndex
	for j < len(rawFileContents) {
		if !unicode.IsDigit(rune(rawFileContents[j])) && rawFileContents[j] != '.' {
			auxNumberList = append(auxNumberList, numberLiteral)
			numberLiteral = ""
			break
		}
		if rawFileContents[j] == '.' {
			auxNumberList = append(auxNumberList, numberLiteral)
			numberLiteral = ""
			auxNumberList = append(auxNumberList, string(rawFileContents[j]))
		}
		if unicode.IsDigit(rune(rawFileContents[j])) {
			numberLiteral += string(rawFileContents[j])
		}
		j++
	}
	if numberLiteral != "" {
		auxNumberList = append(auxNumberList, numberLiteral)
	}
	return auxNumberList, j
}

func handlePeekEqual(currChar rune, peekChar rune, index *int, tokens *[]Token, singleToken TokenType, doubleToken TokenType) {
	if peekChar == '=' {
		*tokens = append(*tokens, Token{Type: doubleToken, StringValue: string(currChar) + string(peekChar), Value: "null"})
		*index += 1
	} else {
		*tokens = append(*tokens, Token{Type: singleToken, StringValue: string(currChar), Value: "null"})
	}
}
