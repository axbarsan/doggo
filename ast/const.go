package ast

import (
	"bytes"

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

func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")

	if cs.Value != nil {
		out.WriteString(cs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
