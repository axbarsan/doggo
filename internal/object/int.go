package object

import (
	"fmt"
)

const (
	INTEGER_OBJ = "INTEGER"
)

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) MapKey() MapKey {
	mk := MapKey{
		Type:  i.Type(),
		Value: uint64(i.Value),
	}

	return mk
}
