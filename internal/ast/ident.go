package ast

import (
	"github.com/axbarsan/doggo/internal/token"
)

type Identifier struct {
	Token token.Token // Any identifier.
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
