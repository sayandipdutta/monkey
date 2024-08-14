package ast

import (
	"bytes"

	"github.com/sayandipdutta/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (prog *Program) String() string {
	var out bytes.Buffer

	for _, s := range prog.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (stmt *LetStatment) statementNode()       {}
func (stmt *LetStatment) TokenLiteral() string { return stmt.Token.Literal }
func (stmt *LetStatment) String() string {
	var out bytes.Buffer

	out.WriteString(stmt.TokenLiteral() + " ")
	out.WriteString(stmt.Name.TokenLiteral() + " ")
	out.WriteString("= ")
	if stmt.Value != nil {
		out.WriteString(stmt.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (ident *Identifier) expressionNode()      {}
func (ident *Identifier) TokenLiteral() string { return ident.Token.Literal }
func (ident *Identifier) String() string       { return ident.Value }

// ReturnStatement
type ReturnStatement struct {
	Value Expression
	Token token.Token
}

func (stmt *ReturnStatement) statementNode()       {}
func (stmt *ReturnStatement) TokenLiteral() string { return stmt.Token.Literal }
func (stmt *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(stmt.TokenLiteral())
	if stmt.Value != nil {
		out.WriteString(stmt.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement
type ExpressionStatement struct {
	Expression Expression
	Token      token.Token
}

func (stmt *ExpressionStatement) statementNode()       {}
func (stmt *ExpressionStatement) TokenLiteral() string { return stmt.Token.Literal }
func (stmt *ExpressionStatement) String() string {
	var out bytes.Buffer

	if stmt.Expression != nil {
		out.WriteString(stmt.Expression.String())
	}
	out.WriteString(";")
	return out.String()
}
