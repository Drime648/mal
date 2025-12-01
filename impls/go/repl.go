package main

import "fmt"

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
