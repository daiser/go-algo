package main

import (
	"fmt"

	"git.lan.316.su/go/algo/shared"
)

func isFound(a [][]int, indexes []int) bool {
	found := true
	for i := range indexes {
		found = found && (a[0][indexes[0]] == a[i][indexes[i]])
	}
	return found
}

func main() {
	// http://algolist.manual.ru/olimp/poi_prb.php
	a := [][]int{
		{0, 0, 3},
		{3, 3, 3},
		{1, 2, 3},
		{3, 4, 5},
		{-1000, 3, 1000},
	}

	n, m := shared.GetSize2D(a)
	indexes := make([]int, n)
	for indexes[0] < m && !isFound(a, indexes) {
		indexes[0]++
		for i := 1; i < n; i++ {
			for ; indexes[i] < m && a[i][indexes[i]] < a[0][indexes[0]]; indexes[i]++ {
			}
		}
	}
	if indexes[0] < m {
		println(a[0][indexes[0]])
		fmt.Printf("%v\n", indexes)
	}
}
