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

func read(line string) string {
	return line
}
