package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens
	// set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

