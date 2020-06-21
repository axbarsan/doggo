package ast

import (
	"bytes"
	"strings"

	"github.com/axbarsan/doggo/internal/token"
)

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token.
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	var out bytes.Buffer
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}
