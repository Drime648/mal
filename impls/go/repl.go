package main

import "fmt"

type (
	envFunc  func(args ...SExpr) SExpr
	listpEnv map[string]envFunc
)

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

func eval(input SExpr, env listpEnv) (any, error) {
	switch input.typ {
	case SExprAtom:
		switch input.atom.typ {
		case AtomNumber:
			return input, nil
		case AtomSymbol:
			envFunc, exists := env[input.atom.symbol.data]
			if !exists {
				return nil, fmt.Errorf("undefined symbol: %s", input.atom.symbol.data)
			}
			return envFunc, nil
		}
	case SExprList:
		var lamdbaFunc envFunc
		args := []SExpr{}
		for idx, sExpr := range input.list {
			if idx == 0 {
				f, err := eval(sExpr, env)
				if err != nil {
					return SExpr{}, err
				}
				switch f.(type) {
				case SExpr:
					return SExpr{}, fmt.Errorf("first element in list is not a function")
				case envFunc:
					lamdbaFunc = f.(envFunc)
				}
			}
		}
		return lamdbaFunc(args...), nil
	}
	return input, nil
}

func Rep(input string) {
	sExpr, err := readStr(input)
	if err != nil {
		return
	}
	sExpr, err = eval(sExpr, replEnv)
	if err != nil {
		fmt.Println(err)
	}
	res := printLisp(sExpr)
	fmt.Println(res)
}
