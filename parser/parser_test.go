package parser

import (
	"testing"
	"monkey/ast"
	"monkey/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")		
	}
	if len(program.Statments) != 3 {
		t.Fatalf("program.Statments does not contain 3 statements. got=%d", len(program.Statements))
	}

	test := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !TestLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}