package main

import (
	"fmt"
	"github.com/kingmidas74/leetcode/internal/problem/roman_to_integer"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments")
	}

	sol := roman_to_integer.New()
	result := sol.Solve(os.Args[1])

	fmt.Print(result)
}
