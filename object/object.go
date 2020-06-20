package object

type Type string

// Object is the internal representation of any type in the doggo language.
type Object interface {
	Type() Type
	Inspect() string
}
