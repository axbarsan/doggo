package ast

import (
	"github.com/axbarsan/doggo/internal/token"
)

type IntegerLiteral struct {
	Token token.Token // Any integer.
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
