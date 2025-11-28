package main

type SExprType int

const (
	SExprAtom SExprType = iota
	SExprList
)

type SExpr struct {
	typ  SExprType
	list []SExpr
	atom Atom
}

type AtomType int

const (
	AtomSymbol AtomType = iota
	AtomNumber
)

type Atom struct {
	typ    AtomType
	symbol Symbol
	number Number
}

type Symbol struct {
	data string
}

type Number struct {
	data float64
}
