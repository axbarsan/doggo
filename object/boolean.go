package object

import (
	"fmt"
)

const (
	BOOLEAN_OBJ = "BOOLEAN"
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
