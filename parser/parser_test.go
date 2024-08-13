package parser_test

import (
	"testing"

	"github.com/sayandipdutta/monkey/ast"
	"github.com/sayandipdutta/monkey/lexer"
	"github.com/sayandipdutta/monkey/parser"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements, got (%d)", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, expected string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("mismatched token literal, expected='let', got=%s", stmt.TokenLiteral())
		return false
	}

	letstmt, ok := stmt.(*ast.LetStatment)
	if !ok {
		t.Errorf("Not a Let Statement. got=%T", stmt)
		return false
	}

	if letstmt.Name.Value != expected {
		t.Errorf("mismatched value. expected=%s, got=%s", expected, letstmt.Name.Value)
		return false
	}

	if letstmt.Name.TokenLiteral() != expected {
		t.Errorf("mismatched name token literal. expected=%s, got=%s", expected, letstmt.Name.TokenLiteral())
		return false
	}

	return true
}
