package main

import (
	"strconv"
)

func printLisp(sExpr SExpr) string {
	switch sExpr.typ {
	case SExprList:
		return printList(sExpr.list)
	case SExprAtom:
		return printAtom(sExpr.atom)
	}
	return ""
}

func printList(list []SExpr) string {
	result := "("
	for i, sExpr := range list {
		if i > 0 {
			result += " " // space separator
		}
		result += printLisp(sExpr)
	}
	result += ")"
	return result
}

func printAtom(atom Atom) string {
	switch atom.typ {
	case AtomNumber:
		return strconv.FormatFloat(atom.number.data, 'g', -1, 64)
	case AtomSymbol:
		return atom.symbol.data
	}
	return ""
}
