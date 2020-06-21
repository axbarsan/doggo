package ast

import (
	"bytes"

	"github.com/axbarsan/doggo/internal/token"
)

type PrefixExpression struct {
	Token    token.Token // A prefix token (e.g. 'token.BANG').
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
