package ast

import (
	"bytes"
	"fmt"

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

// IntegerLiteral
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (ident *IntegerLiteral) expressionNode()      {}
func (ident *IntegerLiteral) String() string       { return fmt.Sprintf("%d", ident.Value) }
func (ident *IntegerLiteral) TokenLiteral() string { return ident.Token.Literal }

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
	return out.String()
}

// IntegerExpression
type IntegerExpression struct {
	Expression Expression
	Token      token.Token
}

func (stmt *IntegerExpression) expressionNode()      {}
func (stmt *IntegerExpression) TokenLiteral() string { return stmt.Token.Literal }
func (stmt *IntegerExpression) String() string       { return stmt.TokenLiteral() }

// PrefixExpression
type PrefixExpression struct {
	Right    Expression
	Operator string
	Token    token.Token
}

func (expr *PrefixExpression) expressionNode()      {}
func (expr *PrefixExpression) TokenLiteral() string { return expr.Token.Literal }
func (expr *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", expr.Operator, expr.Right.String())
}

type InfixExpression struct {
	Left     Expression
	Right    Expression
	Token    token.Token
	Operator string
}

func (expr *InfixExpression) expressionNode()      {}
func (expr *InfixExpression) TokenLiteral() string { return expr.Token.Literal }
func (expr *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", expr.Left.String(), expr.Operator, expr.Right.String())
}
