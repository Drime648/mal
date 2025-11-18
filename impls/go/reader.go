package main

import (
	"fmt"
	"regexp"
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
	if err != nil {
		return nil, err
	}
	if firstToken == "(" {
		sExpr, err := r.readList()
		if err != nil {
			return nil, err
		}
		return sExpr, nil
	} else {
		sExpr, err := r.readAtom()
		if err != nil {
			return nil, err
		}
		return sExpr, nil
	}
}

func (r *Reader) readList() (SExpr, error) {
	list := List{}
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
			list.values = append(list.values, sExpr)
		}
	}
	return list, nil
}

func (r *Reader) readAtom() (SExpr, error) {
	return nil, nil
}

func read(line string) string {
	return line
}
