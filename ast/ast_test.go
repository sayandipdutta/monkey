package ast_test

import (
	"testing"

	"github.com/sayandipdutta/monkey/ast"
	"github.com/sayandipdutta/monkey/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatment{
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Token: token.Token{Type: token.LET, Literal: "let"},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Fatalf("String mismatch. Expected %s, got %s", "let myVar = anotherVar", program.String())
	}
}
