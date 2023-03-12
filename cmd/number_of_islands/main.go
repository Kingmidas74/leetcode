package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kingmidas74/leetcode/internal/problem/number_of_islands"
)

func main() {
	if len(os.Args) < 4 {
		panic("Not enough arguments")
	}

	rowCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	colCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	grid := make([]string, rowCount*colCount)
	for i := 3; i < len(os.Args); i++ {
		grid[i-3] = os.Args[i]
	}

	sol := number_of_islands.New()
	result, err := sol.Solve(grid, rowCount, colCount)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)
}
