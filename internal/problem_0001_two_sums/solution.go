package problem_0001_two_sums

/*
Solve
Description: Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
*/
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
