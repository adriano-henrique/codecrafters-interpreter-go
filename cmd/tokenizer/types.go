package tokenizer

import (
	"fmt"
)

type ErrorType int

const (
	UNEXPECTED_CHARACTER ErrorType = iota
	UNTERMINATED_STRING
)

type Error struct {
	Type  ErrorType
	Value string
	Line  int
}

type TokenType int

const (
	EOF TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	STAR
	SLASH
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	STRING
)

type Token struct {
	Type        TokenType
	StringValue string
	Value       string
}

func (e Error) String() string {
	switch e.Type {
	case UNEXPECTED_CHARACTER:
		return fmt.Sprintf("[line %d] Error: Unexpected character: %s", e.Line, e.Value)
	case UNTERMINATED_STRING:
		return fmt.Sprintf("[line %d] Error: Unterminated string.", e.Line)
	default:
		return "Unknown error"
	}
}

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case MINUS:
		return "MINUS"
	case PLUS:
		return "PLUS"
	case SEMICOLON:
		return "SEMICOLON"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case STRING:
		return "STRING"
	default:
		return "UNKNOWN"
	}
}
