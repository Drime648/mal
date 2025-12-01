package main

func makeNumber(data float64) SExpr {
	return SExpr{
		typ: SExprAtom,
		atom: Atom{
			typ:    AtomNumber,
			number: Number{data: data},
		},
	}
}
