package parser

import (
	"testing"

	"github.com/axbarsan/doggo/ast"
	"github.com/axbarsan/doggo/lexer"
)

func TestConstStatements(t *testing.T) {
	input := `
const x = 5;
const y = 10;
const foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
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
		if !HandleTestConstStatement(t)(stmt, tc.expectedIdentifier) {
			return
		}
	}
}

func HandleTestConstStatement(t *testing.T) func(ast.Statement, string) bool {
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
