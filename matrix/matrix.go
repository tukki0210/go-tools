package matrix

import (
	"fmt"
)

func matrixTest() {
	// MakeMatrix関数の動作確認
	matrix := MakeMatrix(3, 4)
	PrintMatrix(matrix)

	// addMatrix関数の動作確認
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	b := [][]int{{7, 8, 9}, {10, 11, 12}}
	c := AddMatrix(a, b)
	PrintMatrix(c)
}


// 2次元スライスの作成
func MakeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// 2つの2次元スライスの足し算
func AddMatrix(a, b [][]int) [][]int {
	rows := len(a)
	cols := len(a[0])
	c := MakeMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

// 2次元スライスの表示
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}