package main
import (
	"fmt"
	"sort"
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
	matrix := makeMatrix(3, 4)
	printMatrix(matrix)

	// addMatrix関数の動作確認
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	b := [][]int{{7, 8, 9}, {10, 11, 12}}
	c := addMatrix(a, b)
	printMatrix(c)

	// quickSort関数の動作確認
	nums = []int{3, 2, 1, 4, 5}
	fmt.Println(quickSort(nums))

	// binarySearch関数の動作確認
	nums = []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(nums, 3))
	fmt.Println(binarySearch(nums, 6))

	// fib関数の動作確認
	fmt.Println(fib(10))

	// knapsack関数の動作確認
	weights := []int{2, 1, 3}
	values := []int{3, 2, 4}
	maxWeight := 5
	fmt.Println(knapsack(weights, values, maxWeight))


	// sort関数の動作確認
	nums = []int{3, 2, 1, 4, 5}
	sort.Ints(nums)
	
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

// 二分探索
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// 動的計画法(フィボナッチ数列)
func fib(n int) int {
    dp := make([]int, n+1)
    dp[0] = 0
    dp[1] = 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

// 動的計画法(ナップサック問題)
func knapsack(weights []int, values []int, maxWeight int) int {
    n := len(weights)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, maxWeight+1)
    }
    
    for i := 1; i <= n; i++ {
        for w := 1; w <= maxWeight; w++ {
            if weights[i-1] <= w {
                dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
            } else {
                dp[i][w] = dp[i-1][w]
            }
        }
    }
    return dp[n][maxWeight]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 深さ優先探索
func dfs() {
	// 隣接リスト
	graph := [][]int{
		{1, 2},
		{0, 3, 4},
		{0, 5, 6},
		{1},
		{1},
		{2},
		{2},
	}
	// 訪問済みの頂点
	seen := make([]bool, len(graph))
	// 深さ優先探索
	var dfs func(v int)
	
	dfs = func(v int) {
		seen[v] = true
		for _, next := range graph[v] {
			if seen[next] {
				continue
			}
			dfs(next)
		}
	}
	dfs(0)
}

// PairList型の定義
type PairList [][]int

// sort.Interfaceを実装
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i][0] < p[j][0] } // 1つ目の要素を基に比較
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
