// lexer/lexer_test.go

package lexer

import (
	"testing"
	"monkey/token"
)

type Lexer struct {
	input		string
	position	int 	// current position in input, ie: current char
	readPosition	int 	// current reading position in input, ie: after current char
	ch		byte	// current char being examined; note - bc we're using a byte, we are only supporting ASCII
}

func New(input string) *Lexer {
	l := &Lexer(input: input)
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += l
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	
	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral	string
}{
	{token.ASSIGN, "="},
	{token.PLUS, "+"},
	{token.LPAREN, "("},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.RBRACE, "}"},
	{token.COMMA, ","},
	{token.SEMICOLON, ";"},
	{token.EOF, ""},
}

l := New(input)

for i, tt := range tests {
	tok := l.NextToken()

	if tok.Type != tt.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			i, tt.expectedType, tok.Type)
		}

	if tok.Literal != tt.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			i, tt.expectedLiteral, tok.Literal)
		}
	}
}
