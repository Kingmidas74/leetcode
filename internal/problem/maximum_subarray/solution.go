package maximum_subarray

import "github.com/kingmidas74/leetcode/pkg/math_ext"

// Solve for https://leetcode.com/problems/maximum-subarray/
func (s *Solution) Solve(nums []int) int {
	l := nums[0]
	r := nums[0]

	for i := 1; i < len(nums); i++ {
		r = math_ext.Max(nums[i], r+nums[i])
		l = math_ext.Max(l, r)
	}

	return l
}
