package ast

import (
	"bytes"

	"github.com/axbarsan/doggo/token"
)

type InfixExpression struct {
	Token    token.Token // The operator token (e.g. '+').
	Operator string
	Left     Expression
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
