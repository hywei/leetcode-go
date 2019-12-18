package p1

import (
	"fmt"
	"sort")

type INum struct {
	value int
	idx int
}

type ByValue []INum

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func twoSum(nums []int, target int) []int {

	inums := make([]INum, len(nums))
	for idx, value := range nums {
		inums[idx].value = value
		inums[idx].idx = idx
	}

	r := make([]int, 2)
	sort.Sort(ByValue(inums))
	fmt.Println(inums)
	for idx1, in := range inums {
		remain := target - in.value
		inremain := inums[idx1+1:]
		idx2 := sort.Search(len(inremain), func(i int) bool { return inremain[i].value >= remain })
		//fmt.Println(remain, inremain, idx2)
		if idx2 < len(inremain) && inremain[idx2].value == remain {
			r[0], r[1] = inums[idx1].idx, inremain[idx2].idx
			return r
		}
	}
	return r
}
