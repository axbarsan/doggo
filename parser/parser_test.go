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
	testCases := []struct {
		input              string
		expectedIdentifier string
	}{
		{"const x = 5;", "x"},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t)(p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testConstStatement(t)(stmt, tc.expectedIdentifier) {
			return
		}
	}
}

func testConstStatement(t *testing.T) func(ast.Statement, string) bool {
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
	testCases := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return foobar;", "foobar"},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t)(p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
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
		t.Fatalf("program doesn't have enough statements. got=%d", len(program.Statements))
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
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!foobar;", "!", "foobar"},
		{"-foobar;", "-", "foobar"},
		{"!true;", "!", true},
		{"!false;", "!", false},
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
		if !testLiteralExpression(t)(exp.Right, tc.value) {
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

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tc := range infixTests {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t)(p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		if !testInfixExpression(t)(stmt.Expression, tc.leftValue, tc.operator, tc.rightValue) {
			return
		}
	}
}

func testIdentifier(t *testing.T) func(ast.Expression, string) bool {
	return func(exp ast.Expression, value string) bool {
		ident, ok := exp.(*ast.Identifier)
		if !ok {
			t.Errorf("exp not *ast.Identifier. got=%T", exp)

			return false
		}

		if ident.Value != value {
			t.Errorf("ident.Value not %s. got=%s", value, ident.Value)

			return false
		}

		if ident.TokenLiteral() != value {
			t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())

			return false
		}

		return true
	}
}

func testLiteralExpression(t *testing.T) func(ast.Expression, interface{}) bool {
	return func(exp ast.Expression, expected interface{}) bool {
		switch v := expected.(type) {
		case int:
			return testIntegerLiteral(t)(exp, int64(v))
		case int64:
			return testIntegerLiteral(t)(exp, v)
		case string:
			return testIdentifier(t)(exp, v)
		case bool:
			return testBooleanLiteral(t)(exp, v)
		}
		t.Errorf("type of exp not handled. got=%T", exp)

		return false
	}
}

func testInfixExpression(t *testing.T) func(ast.Expression, interface{}, string, interface{}) bool {
	return func(exp ast.Expression, left interface{}, operator string, right interface{}) bool {
		opExp, ok := exp.(*ast.InfixExpression)
		if !ok {
			t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)

			return false
		}

		if !testLiteralExpression(t)(opExp.Left, left) {
			return false
		}

		if opExp.Operator != operator {
			t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)

			return false
		}

		if !testLiteralExpression(t)(opExp.Right, right) {
			return false
		}

		return true
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	testCases := []struct {
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
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t)(p)

		actual := program.String()
		if actual != tc.expected {
			t.Errorf("expected=%q, got=%q", tc.expected, actual)
		}
	}
}

func TestBooleanExpression(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t)(p)

	if len(program.Statements) != 1 {
		t.Fatalf("program doesn't have enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Boolean)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != true {
		t.Errorf("ident.Value not %s. got=%t", "true", ident.Value)
	}
	if ident.TokenLiteral() != "true" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "true", ident.TokenLiteral())
	}
}

func testBooleanLiteral(t *testing.T) func(ast.Expression, bool) bool {
	return func(il ast.Expression, value bool) bool {
		boolean, ok := il.(*ast.Boolean)
		if !ok {
			t.Errorf("il not *ast.Boolean. got=%T", il)
			return false
		}

		if boolean.Value != value {
			t.Errorf("integer.Value not %t. got %t", value, boolean.Value)
			return false
		}

		if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
			t.Errorf("integer.TokenLiteral not %t. got=%s", value, boolean.TokenLiteral())
			return false
		}

		return true
	}
}
