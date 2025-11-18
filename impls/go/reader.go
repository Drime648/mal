package main

import "fmt"

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

func read(line string) string {
	return line
}
