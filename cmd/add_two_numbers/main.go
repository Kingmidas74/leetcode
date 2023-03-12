package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kingmidas74/leetcode/internal/problem/add_two_numbers"
)

func main() {
	if len(os.Args) < 4 {
		panic("Not enough arguments")
	}

	arrLen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	firstArr := make([]int, arrLen)
	for i := 2; i < arrLen+2; i++ {
		firstArr[i-2], err = strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}
	}

	secondArr := make([]int, arrLen)
	for i := arrLen + 2; i < arrLen*2+2; i++ {
		secondArr[i-arrLen-2], err = strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}
	}

	sol := add_two_numbers.New()
	result := sol.Solve(firstArr, secondArr)

	for _, num := range result {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
