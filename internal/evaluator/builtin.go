package evaluator

import (
	"github.com/axbarsan/doggo/internal/object"
)

var builtin = map[string]*object.Builtin{
	"len": {
		Fn: lenFn,
	},
}

func lenFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}

	default:
		return newError("argument to 'len' is not supported, got %s", args[0].Type())
	}
}