package main

import (
	"fmt"
	"sort"
)

// Int64Slice attaches the methods of Interface to []int64, sorting in increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int64Slice) Sort() { sort.Sort(p) }

func main() {
	v := []int64{39121213123, 8912123131231, 4, 82123, 5, 2, 1}
	a := Int64Slice(v)
	fmt.Println(sort.IsSorted(a))
	if !sort.IsSorted(a) {
		sort.Sort(a)
	}
	fmt.Println(sort.IsSorted(a))
	fmt.Println(a)
	//ids := []int64{39121213123, 8912123131231, 4, 82123, 5, 2, 1}

	//// 昇順ソート
	//sort.Slice(ids, func(i, j int) bool {
	//	return ids[i] < ids[j]
	//})

	//// [1 2 3 4 5]
	//fmt.Println(ids)
}
