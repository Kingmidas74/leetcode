package number_of_islands

import "strconv"

func (s *Solution) Solve(grid []string, rowCount, colCount int) (int, error) {
	gridBytes := make([][]byte, rowCount)
	for i := 0; i < rowCount; i++ {
		gridBytes[i] = make([]byte, colCount)
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			cv := grid[i*colCount+j]
			v, err := strconv.ParseBool(cv)
			if err != nil {
				return 0, err
			}
			//TODO: maybe replace to [][]bool
			if v {
				gridBytes[i][j] = 1
			} else {
				gridBytes[i][j] = 0
			}
		}
	}

	return s.solve(gridBytes), nil
}

// Solve for https://leetcode.com/problems/number-of-islands/
func (s *Solution) solve(grid [][]byte) int {
	rowCount := len(grid)
	if rowCount == 0 {
		return 0
	}
	colCount := len(grid[0])
	if colCount == 0 {
		return 0
	}

	islandsCount := 0
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			if grid[row][col] == 1 {
				islandsCount++
				s.dfs(grid, rowCount, colCount, row, col)
			}
		}
	}

	return islandsCount
}

func (s *Solution) dfs(grid [][]byte, rowCount, colCount, currentRow, currentCol int) {
	grid[currentRow][currentCol] = 0

	if currentRow-1 >= 0 && grid[currentRow-1][currentCol] == 1 {
		s.dfs(grid, rowCount, colCount, currentRow-1, currentCol)
	}
	if currentRow+1 < rowCount && grid[currentRow+1][currentCol] == 1 {
		s.dfs(grid, rowCount, colCount, currentRow+1, currentCol)
	}
	if currentCol-1 >= 0 && grid[currentRow][currentCol-1] == 1 {
		s.dfs(grid, rowCount, colCount, currentRow, currentCol-1)
	}
	if currentCol+1 < colCount && grid[currentRow][currentCol+1] == 1 {
		s.dfs(grid, rowCount, colCount, currentRow, currentCol+1)
	}
}
