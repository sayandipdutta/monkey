package parser

import (
	"github.com/sayandipdutta/monkey/ast"
	"github.com/sayandipdutta/monkey/lexer"
	"github.com/sayandipdutta/monkey/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
