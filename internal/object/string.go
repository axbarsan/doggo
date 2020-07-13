package object

import (
	"hash/fnv"
)

const (
	STRING_OBJ = "STRING"
)

type String struct {
	Value string
}

func (s *String) Type() Type {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) MapKey() MapKey {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s.Value))

	mk := MapKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}

	return mk
}
