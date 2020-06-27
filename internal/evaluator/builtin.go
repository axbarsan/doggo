package evaluator

import (
	"github.com/axbarsan/doggo/internal/object"
)

var builtin = map[string]*object.Builtin{
	"length": {
		Fn: lengthFn,
	},
	"lastIndex": {
		Fn: lastIndexFn,
	},
	"tail": {
		Fn: tailFn,
	},
}

func lengthFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}

	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}

	default:
		return newError("argument to 'length' is not supported, got %s", args[0].Type())
	}
}

func lastIndexFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'lastIndex' must be of type ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return &object.Integer{Value: int64(len(arr.Elements) - 1)}
	}

	return NULL
}

func tailFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'tail' must be of type ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	var newElements []object.Object

	if length > 0 {
		newElements = make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
	}

	return &object.Array{Elements: newElements}
}
