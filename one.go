package one

type One = interface{}
type Func = interface{}

type BaseFunctions interface {
	Equals(other One) bool
	HashCode() int
	String() string
}
