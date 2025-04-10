package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	grids := [][]byte{
		{}, {}, {},
		{}, {}, {},
		{}, {}, {},
	}

	gridSquareSize := 3
	gridValidation := 1
	rowCount := 0

	for colCount, row := range board {
		if colCount < 3 {
			gridValidation = 0
		} else if colCount < 6 {
			gridValidation = 3
		} else {
			gridValidation = 6
		}

		// Step 1 - Validate row.
		rowMapCount := make(map[byte]int)
		for _, rowIndexValue := range row {
			if rowCount%gridSquareSize == 0 {
				gridValidation++
			}

			if rowCount == 0 {
				gridValidation = 1
			}

			grids[gridValidation-1] = append(grids[gridValidation-1], rowIndexValue)

			rowCount++
			if rowIndexValue == 46 {
				continue
			}

			if _, ok := rowMapCount[rowIndexValue]; ok {
				return false
			}

			rowMapCount[rowIndexValue]++
		}

		// Step 2 - Validate column
		for j, _ := range row {
			columnMapCount := make(map[byte]int)

			for c := 0; c < len(board); c++ {
				colIndexValue := board[c][j]

				if colIndexValue == 46 {
					continue
				}

				if _, ok := columnMapCount[colIndexValue]; ok {
					return false
				}

				columnMapCount[colIndexValue]++
			}
		}
	}

	// Step 3 - Validate Grid
	for i, grid := range grids {
		gridMapCount := make(map[byte]int)

		for j := 0; j < len(grid); j++ {
			gridIndexValue := grids[i][j]

			if gridIndexValue == 46 {
				continue
			}

			if _, ok := gridMapCount[gridIndexValue]; ok {
				return false
			}

			gridMapCount[gridIndexValue]++
		}
	}
	return true
}

func main() {
	// Success Case
	arr := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// Error case
	// arr := [][]byte{
	// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

	// arr := [][]byte{
	// 	{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
	// 	{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
	// 	{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
	// 	{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
	// 	{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
	// 	{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	// }

	fmt.Print(isValidSudoku(arr))
}
