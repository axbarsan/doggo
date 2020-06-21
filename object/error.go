package object

import (
	"fmt"
)

const (
	ERROR_OBJ = "ERROR"
)

type Error struct {
	Message string
}

func (e *Error) Type() Type {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return fmt.Sprintf("ERROR: %s", e.Message)
}
