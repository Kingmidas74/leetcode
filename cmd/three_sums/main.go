package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kingmidas74/leetcode/internal/problem/three_sums"
)

func main() {
	if len(os.Args) < 5 {
		panic("Not enough arguments")
	}

	requiredSum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	nums := make([]int, len(os.Args)-1)
	for i := 2; i < len(os.Args); i++ {
		nums[i-2], err = strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}
	}

	sol := three_sums.New()
	result := sol.Solve(nums, requiredSum)

	for _, num := range result {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
