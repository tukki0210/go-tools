package algo

import (
	"fmt"
	"sort"
)

func sortTest() {

	// quickSort関数の動作確認
	nums := []int{3, 2, 1, 4, 5}
	fmt.Println(QuickSort(nums))

	pairs := PairList{
		{2, 3},
		{1, 2},
		{3, 3},
		{4, 5},
	}

	// 指定した要素でソート
	sort.Sort(pairs)
	fmt.Println(pairs)
}

// クイックソート
func QuickSort(nums []int) []int {
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
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}


// PairList型の定義
type PairList [][]int

// sort.Interfaceを実装
// sort.Sort()を使うために必要なもの
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i][0] < p[j][0] } // 1つ目の要素を基に比較
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
