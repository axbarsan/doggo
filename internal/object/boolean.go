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

func (b *Boolean) MapKey() MapKey {
	var value uint64
	if b.Value {
		value = 1
	}

	mk := MapKey{
		Type:  b.Type(),
		Value: value,
	}

	return mk
}
