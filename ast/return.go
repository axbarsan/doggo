package ast

import (
	"github.com/axbarsan/doggo/token"
)

type ReturnStatement struct {
	Token       token.Token // The token.RETURN token.
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
