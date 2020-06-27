package ast

import (
	"github.com/axbarsan/doggo/internal/token"
)

type StringLiteral struct {
	Token token.Token // The 'token.STRING' token.
	Value string
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}
