package ast

import "github.com/sayandipdutta/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program
type Program struct {
	Statements []Statement
}

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatment
type LetStatment struct {
	Name  *Identifier
	Value Expression
	Token token.Token
}

func (stmt *LetStatment) statementNode() {}
func (stmt *LetStatment) TokenLiteral() string {
	return stmt.Token.Literal
}

// Identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (ident *Identifier) expressionNode() {}
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}

// LetStatment
type ReturnStatement struct {
	Value Expression
	Token token.Token
}

func (stmt *ReturnStatement) statementNode() {}
func (stmt *ReturnStatement) TokenLiteral() string {
	return stmt.Token.Literal
}
