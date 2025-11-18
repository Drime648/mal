package main

type Symbol struct {
	sym string
}

type AtomType string

const (
	AtomSymbol AtomType = "symbol"
	AtomNumber AtomType = "number"
)

type Atom struct {
	typ    AtomType
	symbol Symbol
	number float64
}

type SexprType string

const (
	SexprList SexprType = "list"
	SexprAtom SexprType = "atom"
)

type Sexpr struct {
	typ  SexprType
	list []Sexpr
	atom Atom
}
