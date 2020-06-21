package object

import (
	"bytes"
	"strings"

	"github.com/axbarsan/doggo/internal/ast"
)

const (
	FUNCTION_OBJ = "FUNCTION"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	var out bytes.Buffer
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
