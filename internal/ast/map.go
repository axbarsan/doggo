package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/axbarsan/doggo/internal/token"
)

type MapLiteral struct {
	Token token.Token // The 'token.LBRACE' token.
	Pairs map[Expression]Expression
}

func (ml *MapLiteral) expressionNode() {}

func (ml *MapLiteral) TokenLiteral() string {
	return ml.Token.Literal
}

func (ml *MapLiteral) String() string {
	var out bytes.Buffer

	var pairs []string
	for key, value := range ml.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s:%s", key.String(), value.String()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
