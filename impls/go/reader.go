package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Reader struct {
	tokens   []string
	position int
}

func (r *Reader) Next() (string, error) {
	if r.position >= len(r.tokens) {
		return "", fmt.Errorf("reading finished")
	}
	token := r.tokens[r.position]
	r.position++
	return token, nil
}

func (r *Reader) Peek() (string, error) {
	if r.position >= len(r.tokens) {
		return "", fmt.Errorf("reading finished")
	}
	token := r.tokens[r.position]
	return token, nil
}

func readStr(input string) *Reader {
	tokens := tokenize(input)
	r := &Reader{tokens: tokens, position: 0}
	return r
}

func tokenize(input string) []string {
	r, _ := regexp.Compile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
	tokens := r.FindAllString(input, -1)

	return tokens
}

func (r *Reader) readForm() (SExpr, error) {
	firstToken, err := r.Peek()
	sExpr := SExpr{}
	if err != nil {
		return sExpr, err
	}
	if firstToken == "(" {
		list, err := r.readList()
		if err != nil {
			return sExpr, err
		}
		sExpr.list = list
		sExpr.typ = SExprList
		return sExpr, nil
	} else {
		atom, err := r.readAtom()
		if err != nil {
			return sExpr, err
		}
		sExpr.atom = atom
		sExpr.typ = SExprAtom
		return sExpr, nil
	}
}

func (r *Reader) readList() ([]SExpr, error) {
	list := make([]SExpr, 0)
	for {
		currToken, err := r.Peek()
		if err != nil {
			return nil, err
		}
		if currToken == ")" {
			break
		} else {
			sExpr, err := r.readForm()
			if err != nil {
				return nil, err
			}
			list = append(list, sExpr)
		}
	}
	return list, nil
}

func (r *Reader) readAtom() (Atom, error) {
	atom := Atom{}
	atomStr, err := r.Next()
	if err != nil {
		return atom, err
	}
	num, invalid_number := strconv.ParseFloat(atomStr, 64)
	if invalid_number == nil { // no error, so it is a valid number
		atom.typ = AtomNumber
		atom.number = Number{data: num}
	} else {
		atom.typ = AtomSymbol
		atom.symbol = Symbol{data: atomStr}
	}

	return atom, nil
}

func read(line string) string {
	return line
}
