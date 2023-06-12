package main

import (
	"fmt"
	// "sort"

)

func main() {

	// binarySearch関数の動作確認
	// nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(binarySearch(nums, 3))
	// fmt.Println(binarySearch(nums, 6))

	// fib関数の動作確認
	// fmt.Println(fib(10))

	// knapsack関数の動作確認
	weights := []int{3, 1, 3,2,1,5}
	values := []int{3, 2, 6,1,3,85}
	maxWeight := 9
	fmt.Println(knapsack(weights, values, maxWeight))

	// sort関数の動作確認
	// nums = []int{3, 2, 1, 4, 5}
	// sort.Ints(nums)

}

// 動的計画法(ナップサック問題)
func knapsack(weights []int, values []int, maxWeight int) int {
	// n 石の数
	// maxWeight 重さの上限
	// n+1 * maxWeight+1 の二次元スライスにメモを格納する
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxWeight+1)
	}


	for i := 0; i < n; i++ {
		for w := 0; w < maxWeight; w++ {
			if weights[i] <= w+1 {
				dp[i+1][w+1] = max(dp[i][w+1], dp[i][w-weights[i]+1]+values[i])
			} else {
				dp[i+1][w+1] = dp[i][w+1]
			}
		}
	}
	fmt.Println(dp)
	return dp[n][maxWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 最大公約数
func makeGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return makeGCD(b, a%b)
}

// 最小公倍数
func makeLCM(a, b int) int {
	return a * b / makeGCD(a, b)
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
