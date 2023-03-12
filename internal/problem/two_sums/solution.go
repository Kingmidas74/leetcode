package two_sums

// Solve for https://leetcode.com/problems/two-sum/
func (s *Solution) Solve(nums []int, target int) []int {
	for i, n1 := range nums {
		t := target - n1

		for j, n2 := range nums {
			if j == i {
				continue
			}

			if n2 == t {
				return []int{i, j}
			}
		}
	}

	return make([]int, 0)
}
