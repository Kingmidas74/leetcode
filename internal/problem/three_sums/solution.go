package three_sums

import "sort"

// Solve for https://leetcode.com/problems/3sum/
func (s *Solution) Solve(nums []int, target int) [][]int {
	res := make([][]int, 0)
	l := len(nums)

	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1
		for j < k {
			t := nums[i] + nums[j] + nums[k]
			if t == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if t < target {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
