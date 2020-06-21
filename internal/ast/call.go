package ast

import (
	"bytes"
	"strings"

	"github.com/axbarsan/doggo/internal/token"
)

type CallExpression struct {
	Token     token.Token // The '(' token.
	Function  Expression  // Identifier or FunctionLiteral.
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	var out bytes.Buffer
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
