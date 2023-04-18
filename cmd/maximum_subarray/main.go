package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kingmidas74/leetcode/internal/problem/maximum_subarray"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments")
	}

	var err error
	nums := make([]int, len(os.Args)-1)
	for i := 2; i < len(os.Args); i++ {
		nums[i-2], err = strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}
	}

	sol := maximum_subarray.New()
	result := sol.Solve(nums)

	fmt.Print(result)
}
