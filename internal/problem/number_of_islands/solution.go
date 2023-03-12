package number_of_islands

import "strconv"

func (s *Solution) Solve(grid []string, rowCount, colCount int) (int, error) {
	gridBytes := make([][]bool, rowCount)
	for i := 0; i < rowCount; i++ {
		gridBytes[i] = make([]bool, colCount)
	}

	var err error
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			cv := grid[i*colCount+j]
			gridBytes[i][j], err = strconv.ParseBool(cv)
			if err != nil {
				return 0, err
			}
		}
	}

	return s.solve(gridBytes), nil
}

// Solve for https://leetcode.com/problems/number-of-islands/
func (s *Solution) solve(grid [][]bool) int {
	return 0
}
