package main

import "fmt"

type listpEnv map[string]func(args ...SExpr) SExpr

var replEnv = listpEnv{
	"+": func(args ...SExpr) SExpr {
		return makeNumber(args[0].atom.number.data + args[1].atom.number.data)
	},
	"-": func(args ...SExpr) SExpr {
		return makeNumber(args[0].atom.number.data - args[1].atom.number.data)
	},
	"/": func(args ...SExpr) SExpr {
		return makeNumber(args[0].atom.number.data / args[1].atom.number.data)
	},
	"*": func(args ...SExpr) SExpr {
		return makeNumber(args[0].atom.number.data * args[1].atom.number.data)
	},
}

func eval(input SExpr) SExpr {
	return input
}

func Rep(input string) {
	sExpr, err := readStr(input)
	if err != nil {
		return
	}
	sExpr = eval(sExpr)
	res := printLisp(sExpr)
	fmt.Println(res)
}
