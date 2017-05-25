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
	l.readChar()	
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
