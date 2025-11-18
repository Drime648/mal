package main

// SExpr can either be an atom or a list of SExprs
type SExpr interface {
	isSExpr()
}

type List struct {
	values []SExpr
}

func (List) isSExpr() {}

// Atom is an SExpr, and can either be a Symbol or a Number
type Atom interface {
	SExpr
	isAtom()
}

type Symbol struct {
	data string
}

type Number struct {
	data float64
}

func (Symbol) isAtom() {}
func (Number) isAtom() {} // defines the Number as part of an atom
