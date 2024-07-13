package tokenizer

type TokenType int

const (
	EOF TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
)

type Token struct {
	Type  TokenType
	Value string
}

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	default:
		return "UNKNOWN"
	}
}
