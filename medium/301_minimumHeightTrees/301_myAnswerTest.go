package leetcode

import (
	"fmt"
	"testing"
)

type question301 struct {
	para301
	ans301
}

type para301 struct {
	n		int
	edges	[][]int
}

type ans301 struct {
	ans []int
}

func Test_Problem301(t *testing.T)  {
	qs := []question301{
		question301{
			para301{4, [][]int{{1,0}, {1,2}, {1,3}}},
			ans301{[]int{1}},
			},
		question301{
			para301{6, [][]int{{0,3}, {1,3}, {2,3}, {4,3}, {5,4}}},
			ans301{[]int{3, 4}},
			},
		question301{
			para301{7, [][]int{{0,1}, {1,2}, {1,3}, {2,4}, {3,5}, {4,6}}},
			ans301{[]int{3, 4}},
			},
	}

	fmt.Printf("------------------------Leetcode Problem 301------------------------\n")

	for _, q := range qs {
		_, p := q.ans301, q.para301
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMinHeightTrees(p.n, p.edges))
	}
	fmt.Printf("\n\n\n")
}