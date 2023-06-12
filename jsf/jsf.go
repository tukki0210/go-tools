package jsf

import (
	"fmt"
)

func jsfTest(){
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