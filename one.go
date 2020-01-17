package one

type One = interface{}

type BaseFunctions interface {
	Equals(other One) bool
	HashCode() int
	ToString() string
}
