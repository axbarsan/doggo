package ast

import (
	"github.com/axbarsan/doggo/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type ConstStatement struct {
	Token token.Token // The token.CONST token.
	Name  *Identifier
	Value Expression
}

func (cs *ConstStatement) statementNode() {

}

func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal
}

type Identifier struct {
	Token token.Token // The token.IDENT token.
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
