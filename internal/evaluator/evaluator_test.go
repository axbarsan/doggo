package evaluator

import (
	"testing"

	"github.com/axbarsan/doggo/internal/lexer"
	"github.com/axbarsan/doggo/internal/object"
	"github.com/axbarsan/doggo/internal/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
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
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{`"Hello" == "Hello"`, true},
		{`"Hello" == "hello"`, false},
		{`"Hello" != "Hello"`, false},
		{`"Hello" != "hello"`, true},
		{`"Hello" == "H" + "ello"`, true},
		{`"Hello" == "h" + "ello"`, false},
		{`"Hello" != "h" + "ello"`, true},
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

func TestIfElseExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		integer, ok := tc.expected.(int)
		if ok {
			testIntegerObject(t)(evaluated, int64(integer))
		} else {
			testNullObject(t)(evaluated)
		}
	}
}

func testNullObject(t *testing.T) func(object.Object) bool {
	return func(obj object.Object) bool {
		if obj != NULL {
			t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)

			return false
		}

		return true
	}
}

func TestReturnStatements(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{
			`
if (10 > 1) {
    if (10 > 1) {
        return 10;
    }

    return 1;
}
        `,
			10,
		},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		testIntegerObject(t)(evaluated, tc.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	testCases := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if 10 > 1 {
    if 10 > 1 {
        return true + false;
    }

    return 1;
}
`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
		{
			`{"name": "test"}[fn(x) { x }];`,
			"unusable as map key: FUNCTION",
		},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)

			continue
		}

		if errObj.Message != tc.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tc.expectedMessage, errObj.Message)
		}
	}
}

func TestConstStatements(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"const a = 5; a;", 5},
		{"const a = 5 * 5; a;", 25},
		{"const a = 5; const b = a; b;", 5},
		{"const a = 5; const b = a; const c = a + b + 5; c;", 15},
	}

	for _, tc := range testCases {
		testIntegerObject(t)(testEval(tc.input), tc.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"const identity = fn(x) { x; }; identity(5);", 5},
		{"const identity = fn(x) { return x; }; identity(5);", 5},
		{"const double = fn(x) { x * 2; }; double(5);", 10},
		{"const add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"const add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for _, tc := range testCases {
		testIntegerObject(t)(testEval(tc.input), tc.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
const newAdder = fn(x) {
    fn(y) { x + y };
};

const addTwo = newAdder(2);
addTwo(2);`

	testIntegerObject(t)(testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	testCases := []struct {
		input    string
		expected interface{}
	}{
		{`length("")`, 0},
		{`length("four")`, 4},
		{`length("hello world")`, 11},
		{`length([1, 2, 3])`, 3},
		{`length([])`, 0},
		{`length(1)`, "argument to 'length' is not supported, got INTEGER"},
		{`length("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`lastIndex([1, 2, 3])`, 2},
		{`lastIndex([])`, NULL},
		{`lastIndex([3])`, 0},
		{`lastIndex("")`, "argument to 'lastIndex' must be of type ARRAY, got STRING"},
		{`lastIndex([1, 2, 3], [3, 4, 5])`, "wrong number of arguments. got=2, want=1"},
		{`tail([1, 2, 3])`, []int{2, 3}},
		{`tail([])`, []int{}},
		{`tail("")`, "argument to 'tail' must be of type ARRAY, got STRING"},
		{`tail([1, 2, 3], [3, 4, 5])`, "wrong number of arguments. got=2, want=1"},
		{`push([1, 2, 3], 4)`, []int{1, 2, 3, 4}},
		{`push([], 5)`, []int{5}},
		{`push("", 2)`, "first argument to 'push' must be of type ARRAY, got STRING"},
		{`push([3, 4, 5])`, "wrong number of arguments. got=1, want=2"},
		{`push([3, 4, 5], 5, 6)`, "wrong number of arguments. got=3, want=2"},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)

		switch expected := tc.expected.(type) {
		case int:
			testIntegerObject(t)(evaluated, int64(expected))

		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}

		case *object.Null:
			testNullObject(t)(evaluated)
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t)(result.Elements[0], 1)
	testIntegerObject(t)(result.Elements[1], 4)
	testIntegerObject(t)(result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"const i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"const myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"const myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"const myArray = [1, 2, 3]; const i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		integer, ok := tc.expected.(int)
		if ok {
			testIntegerObject(t)(evaluated, int64(integer))
		} else {
			testNullObject(t)(evaluated)
		}
	}
}

func TestMapLiterals(t *testing.T) {
	input := `const two = "two";
{
	"one": 10 - 9,
	two: 1 + 1,
	"thr" + "ee": 6 / 2,
	4: 4,
	true: 5,
	false: 6
}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Map)
	if !ok {
		t.Fatalf("Eval didn't return Map. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.MapKey]int64{
		(&object.String{Value: "one"}).MapKey():   1,
		(&object.String{Value: "two"}).MapKey():   2,
		(&object.String{Value: "three"}).MapKey(): 3,
		(&object.Integer{Value: 4}).MapKey():      4,
		TRUE.MapKey():                             5,
		FALSE.MapKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Map has wrong number of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t)(pair.Value, expectedValue)
	}
}

func TestMapIndexExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
	}

	for _, tc := range testCases {
		evaluated := testEval(tc.input)
		integer, ok := tc.expected.(int)
		if ok {
			testIntegerObject(t)(evaluated, int64(integer))
		} else {
			testNullObject(t)(evaluated)
		}
	}
}
