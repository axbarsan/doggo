package evaluator

import (
	"testing"

	"github.com/axbarsan/doggo/lexer"
	"github.com/axbarsan/doggo/object"
	"github.com/axbarsan/doggo/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		testIntegerObject(t)(evaluated, tc.expected)
	}
}

func testIntegerObject(t *testing.T) func(object.Object, int64) bool {
	return func(obj object.Object, expected int64) bool {
		result, ok := obj.(*object.Integer)
		if !ok {
			t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)

			return false
		}

		if result.Value != expected {
			t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)

			return false
		}

		return true
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		testBooleanObject(t)(evaluated, tc.expected)
	}
}

func testBooleanObject(t *testing.T) func(object.Object, bool) bool {
	return func(obj object.Object, expected bool) bool {
		result, ok := obj.(*object.Boolean)
		if !ok {
			t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)

			return false
		}

		if result.Value != expected {
			t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)

			return false
		}

		return true
	}
}

func TestBangOperator(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		testBooleanObject(t)(evaluated, tc.expected)
	}
}
