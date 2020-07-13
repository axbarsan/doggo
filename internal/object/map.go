package object

import (
	"bytes"
	"fmt"
	"strings"
)

type MapKey struct {
	Type  Type
	Value uint64
}

type Mappable interface {
	Object
	MapKey() MapKey
}

const (
	MAP_OBJ = "MAP"
)

type MapPair struct {
	Key   Mappable
	Value Object
}

type Map struct {
	Pairs map[MapKey]MapPair
}

func (m *Map) Type() Type {
	return MAP_OBJ
}

func (m *Map) Inspect() string {
	var out bytes.Buffer

	var pairs []string
	for _, pair := range m.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
