package parser

import (
	"fmt"
	"strconv"

	"github.com/sayandipdutta/monkey/ast"
	"github.com/sayandipdutta/monkey/lexer"
	"github.com/sayandipdutta/monkey/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // < or >
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer          *lexer.Lexer
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	currToken      token.Token
	peekToken      token.Token
	Errors         []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  lexer,
		Errors: []string{},
	}

	parser.nextToken()
	parser.nextToken()

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefixFn(token.IDENT, parser.parseIdentifier)
	parser.registerPrefixFn(token.INT, parser.parseIntLiteral)

	return parser
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	case token.RETURN:
		return parser.parseReturnStatement()
	case token.IDENT, token.INT:
		return parser.parseExpressionStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatment {
	letstmt := &ast.LetStatment{Token: parser.currToken}

	// TODO: store parseErrors
	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	letstmt.Name = &ast.Identifier{Token: parser.currToken, Value: parser.currToken.Literal}

	// TODO: store parseErrors
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Parse expressions
	for !parser.currTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return letstmt
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	retstmt := &ast.ReturnStatement{Token: parser.currToken}
	parser.nextToken()

	for !parser.currTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return retstmt
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expst := &ast.ExpressionStatement{Token: parser.currToken}
	expst.Expression = parser.parseExpression(LOWEST)

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return expst
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.currToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: parser.currToken,
		Value: parser.currToken.Literal,
	}
}

func (parser *Parser) parseIntLiteral() ast.Expression {
	value, err := strconv.ParseInt(parser.currToken.Literal, 10, 64)
	if err != nil {
		return nil
	}
	return &ast.IntegerLiteral{
		Token: parser.currToken,
		Value: value,
	}
}

func (parser *Parser) peekError(tok token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s", tok, parser.peekToken.Type)
	parser.Errors = append(parser.Errors, msg)
}

func (parser *Parser) expectPeek(tok token.TokenType) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	}
	parser.peekError(tok)
	return false
}

func (parser *Parser) currTokenIs(tok token.TokenType) bool {
	return parser.currToken.Type == tok
}

func (parser *Parser) peekTokenIs(tok token.TokenType) bool {
	return parser.peekToken.Type == tok
}

func (parser *Parser) registerPrefixFn(tok token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tok] = fn
}

func (parser *Parser) registerInfixFn(tok token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tok] = fn
}
