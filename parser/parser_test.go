package parser_test

import (
	"fmt"
	"testing"

	"github.com/sayandipdutta/monkey/ast"
	"github.com/sayandipdutta/monkey/lexer"
	"github.com/sayandipdutta/monkey/parser"
)

func TestLetStatement(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let foobar = 123456;
  `
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseError(t, p)

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

func TestReturnStatement(t *testing.T) {
	input := `
  return 5;
  return 10;
  return 123456;
  `
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseError(t, p)

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements, got (%d)", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		retstmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("Not a Return Statement. got=%T", retstmt)
		}

		if retstmt.TokenLiteral() != "return" {
			t.Fatalf("Wrong literal. Expected=%s, got=%s", "return", retstmt.TokenLiteral())
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

func checkParseError(t *testing.T, p *parser.Parser) {
	errors := p.Errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has generated %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("Statements expected=%d, got=%d", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statement not an ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier, got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("Value expected=foobar, got=%s", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("Statements expected=%d, got=%d", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statement not an ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression not *ast.IntegerLiteral, got=%T", stmt.Expression)
	}

	if ident.Value != 5 {
		t.Fatalf("Value expected=5, got=%d", ident.Value)
	}

	if ident.TokenLiteral() != "5" {
		t.Fatalf("ident.TokenLiteral not %s. got=%s", "5", ident.TokenLiteral())
	}
}

func TestPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-21;", "-", 1},
	}

	for _, tt := range prefixTests {
		lexer := lexer.New(tt.input)
		parser := parser.New(lexer)
		program := parser.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, found=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Expected ExpressionStatement, found=%T", program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Exprected PrefixExpression, got=%q", stmt.Expression)
		}

		if expr.Operator != tt.operator {
			t.Fatalf("Operator mismatch. Expected=%s, got=%s", tt.operator, expr.Operator)
		}

	}
}

func testIntegerLiteral(t *testing.T, intlit ast.Expression, value int64) bool {
	integer, ok := intlit.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("intlit not *ast.IntegerLiteral. got=%T", intlit)
		return false
	}

	if integer.Value != value {
		t.Errorf("Expected integer=%d, got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("Expected integer=%d, got=%s", value, integer.TokenLiteral())
		return false
	}

	return true
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		operator string
		left     int64
		right    int64
	}{
		{"5 + 5;", "+", 5, 5},
		{"5 - 5;", "-", 5, 5},
		{"5 * 5;", "*", 5, 5},
		{"5 / 5;", "/", 5, 5},
		{"5 > 5;", ">", 5, 5},
		{"5 < 5;", "<", 5, 5},
		{"5 == 5;", "==", 5, 5},
		{"5 != 5;", "!=", 5, 5},
	}

	for _, tt := range infixTests {
		lexer := lexer.New(tt.input)
		parser := parser.New(lexer)

		program := parser.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, found=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Expected ExpressionStatement, found=%T", program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Expected InfixExpression, found=%T", stmt.Expression)
		}

		if expr.Operator != tt.operator {
			t.Fatalf("Expected operator to be=%s, found=%s", tt.operator, expr.Operator)
		}

		if !testIntegerLiteral(t, expr.Left, tt.left) {
			t.Fatalf("Expected left operand=%d, found=%d", tt.left, expr.Left)
		}

		if !testIntegerLiteral(t, expr.Right, tt.right) {
			t.Fatalf("Expected right operand=%d, found=%d", tt.right, expr.Right)
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 + 5",
			"(3 + 4)((-5) + 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"3 + 4 * 5 != 3 * 1 - 4 / 5",
			"((3 + (4 * 5)) != ((3 * 1) - (4 / 5)))",
		},
	}

	for _, tt := range tests {
		lexer := lexer.New(tt.input)
		parser := parser.New(lexer)
		program := parser.ParseProgram()

		checkParseError(t, parser)

		if program.String() != tt.expected {
			t.Fatalf("Expected string=%s, found=%s", tt.expected, program.String())
		}
	}
}
