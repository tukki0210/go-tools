package main
import (
	"fmt"
)


func main() {
	nums := []int{1, 2, 3, 4}

	// Map関数の動作確認
	double := Map(nums, func(n int) int {
		return n * 2
	})
	fmt.Println(double)

	// Fillter関数の動作確認
	odds := Fillter(nums, func(n int) bool {
		return n % 2 == 1
	})
	fmt.Println(odds)

	// makeMatrix関数の動作確認
	matrix := makeMatrix(3, 3)
	printMatrix(matrix)

	// addMatrix関数の動作確認
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	b := [][]int{{7, 8, 9}, {10, 11, 12}}
	c := addMatrix(a, b)
	printMatrix(c)

	// quickSort関数の動作確認
	nums = []int{3, 2, 1, 4, 5}
	fmt.Println(quickSort(nums))

}

// JavaScriptのmap関数
func Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// JavaScriptのfillter関数
func Fillter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 2次元スライスの作成
func makeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// 2つの2次元スライスの足し算
func addMatrix(a, b [][]int) [][]int {
	rows := len(a)
	cols := len(a[0])
	c := makeMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

// 2次元スライスの表示
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// 最大公約数
func makeGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return makeGCD(b, a % b)
}
// 最小公倍数
func makeLCM(a, b int) int {
	return a * b / makeGCD(a, b)
}

// クイックソート
func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, v := range nums[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}