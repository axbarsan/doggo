package parser

import (
	"fmt"
	"testing"

	"github.com/axbarsan/doggo/ast"
	"github.com/axbarsan/doggo/lexer"
)

func checkParserErrors(t *testing.T) func(p *Parser) {
	return func(p *Parser) {
		errors := p.Errors()
		if len(errors) == 0 {
			return
		}

		t.Errorf("parser has %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}
}

func TestConstStatements(t *testing.T) {
	input := `
const x = 5;
const y = 10;
const foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t)(p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil.")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tc := range testCases {
		stmt := program.Statements[i]
		if !handleTestConstStatement(t)(stmt, tc.expectedIdentifier) {
			return
		}
	}
}

func handleTestConstStatement(t *testing.T) func(ast.Statement, string) bool {
	return func(s ast.Statement, name string) bool {
		if s.TokenLiteral() != "const" {
			t.Errorf("s.TokenLiteral not 'const'. got=%q", s.TokenLiteral())

			return false
		}

		constStmt, ok := s.(*ast.ConstStatement)
		if !ok {
			t.Errorf("s not *ast.ConstStatement. got=%T", s)

			return false
		}

		if constStmt.Name.Value != name {
			t.Errorf("constStmt.Name.Value not '%s'. got=%s", name, constStmt.Name.Value)

			return false
		}

		if constStmt.Name.TokenLiteral() != name {
			t.Errorf("constStmt.Name.TokenLiteral() not '%s'. got=%s", name, constStmt.Name.TokenLiteral())

			return false
		}

		return true
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t)(p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statemenmts does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)

			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t)(p)

	if len(program.Statements) != 1 {
		t.Fatalf("program doesn't have enoughs statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t)(p)

	if len(program.Statements) != 1 {
		t.Fatalf("program doesn't have enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statemetns[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got %d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tc := range prefixTests {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t)(p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tc.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tc.operator, exp.Operator)
		}
		if !testIntegerLiteral(t)(exp.Right, tc.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T) func(ast.Expression, int64) bool {
	return func(il ast.Expression, value int64) bool {
		integer, ok := il.(*ast.IntegerLiteral)
		if !ok {
			t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
			return false
		}

		if integer.Value != value {
			t.Errorf("integer.Value not %d. got %d", value, integer.Value)
			return false
		}

		if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
			t.Errorf("integer.TokenLiteral not %d. got=%s", value, integer.TokenLiteral())
			return false
		}

		return true
	}
}
