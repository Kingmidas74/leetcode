package main

import (
	"fmt"
	"os"

	"github.com/kingmidas74/leetcode/internal/problem/valid_parentheses"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments")
	}

	sol := valid_parentheses.New()
	result := sol.Solve(os.Args[1])

	fmt.Print(result)
}
