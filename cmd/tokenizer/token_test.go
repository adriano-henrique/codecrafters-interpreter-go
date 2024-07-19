package tokenizer

import "testing"

func verifyExpectedResults(t *testing.T, tokenList []Token, errorList []Error, endIndex int, expectedTokenList []Token, expectedErrorList []Error, expectedEndIndex int) {
	if len(tokenList) != len(expectedTokenList) {
		t.Errorf("Expected %d tokens, got %d", len(expectedTokenList), len(tokenList))
	}
	for i := 0; i < len(tokenList); i++ {
		if tokenList[i].Type != expectedTokenList[i].Type {
			t.Errorf("Expected token type %s, got %s", expectedTokenList[i].Type.String(), tokenList[i].Type.String())
		}
		if tokenList[i].StringValue != expectedTokenList[i].StringValue {
			t.Errorf("Expected token string value %s, got %s", expectedTokenList[i].StringValue, tokenList[i].StringValue)
		}
		if tokenList[i].Value != expectedTokenList[i].Value {
			t.Errorf("Expected token value %s, got %s", expectedTokenList[i].Value, tokenList[i].Value)
		}
	}
	for i := 0; i < len(errorList); i++ {
		if errorList[i].Type != expectedErrorList[i].Type {
			t.Errorf("Expected error type %s, got %s", expectedErrorList[i].String(), errorList[i].String())
		}
		if errorList[i].Value != expectedErrorList[i].Value {
			t.Errorf("Expected error value %s, got %s", expectedErrorList[i].Value, errorList[i].Value)
		}
		if errorList[i].Line != expectedErrorList[i].Line {
			t.Errorf("Expected error line %d, got %d", expectedErrorList[i].Line, errorList[i].Line)
		}
	}
	if endIndex != expectedEndIndex {
		t.Errorf("Expected end index %d, got %d", expectedEndIndex, endIndex)
	}
}

func TestNumberTokenizer(t *testing.T) {
	regularNumber := "1234"
	expectedTokenList := []Token{{Type: 21, StringValue: "1234", Value: "1234"}}
	expectedErrorList := []Error{}
	expectedEndIndex := len(regularNumber)
	tokenList, errorList, endIndex := handleNumberTokenization(regularNumber, 0)
	verifyExpectedResults(t, tokenList, errorList, endIndex, expectedTokenList, expectedErrorList, expectedEndIndex)
}

func TestWeirdNumberTokenizer(t *testing.T) {
	regularNumber := "1234.1234.1234"
	expectedTokenList := []Token{{Type: 21, StringValue: "1234.1234", Value: "1234.1234"}}
	expectedErrorList := []Error{}
	expectedEndIndex := len(regularNumber)
	tokenList, errorList, endIndex := handleNumberTokenization(regularNumber, 0)
	verifyExpectedResults(t, tokenList, errorList, endIndex, expectedTokenList, expectedErrorList, expectedEndIndex)
}

func TestNumberWithDot(t *testing.T) {
	numberWithDot := "1234."
	expectedTokenList := []Token{{Type: 21, StringValue: "1234.", Value: "1234."}, {Type: 6, StringValue: ".", Value: "null"}}
	expectedErrorList := []Error{}
	expectedEndIndex := len(numberWithDot)
	tokenList, errorList, endIndex := handleNumberTokenization(numberWithDot, 0)
	verifyExpectedResults(t, tokenList, errorList, endIndex, expectedTokenList, expectedErrorList, expectedEndIndex)
}

func TestNumberWithDotAndMoret(t *testing.T) {
	numberWithDot := "1234.123..12.12"
	expectedTokenList := []Token{
		{Type: 21, StringValue: "1234.123", Value: "1234.123"},
		{Type: 6, StringValue: ".", Value: "null"},
		{Type: 6, StringValue: ".", Value: "null"},
		{Type: 21, StringValue: "12.12", Value: "12.12"},
	}
	expectedErrorList := []Error{}
	expectedEndIndex := len(numberWithDot)
	tokenList, errorList, endIndex := handleNumberTokenization(numberWithDot, 0)
	verifyExpectedResults(t, tokenList, errorList, endIndex, expectedTokenList, expectedErrorList, expectedEndIndex)
}
