package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kingmidas74/leetcode/internal/problem_0001_two_sums"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
*/
func main() {
	if len(os.Args) < 4 {
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

	sol := problem_0001_two_sums.New()
	result := sol.Solve(nums, requiredSum)

	for _, num := range result {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
