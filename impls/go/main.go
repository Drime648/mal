package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("user> ")
		if !s.Scan() {
			return
		}
		input := s.Text()
		Rep(input)
	}
}
