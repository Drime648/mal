package main

import "fmt"

func eval(line string) string {
	return line
}

func print(line string) {
	fmt.Println(line)
}

func Rep(input string) {
	x := read(input)
	x = eval(x)
	print(x)
}
