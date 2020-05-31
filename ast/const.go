package ast

import (
	"github.com/axbarsan/doggo/token"
)

type ConstStatement struct {
	Token token.Token // The token.CONST token.
	Name  *Identifier
	Value Expression
}

func (cs *ConstStatement) statementNode() {}

func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal
}
