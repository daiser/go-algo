package main

import "fmt"

func getSize(a [][]int) (n, m int) {
	n = len(a)
	if n < 1 {
		panic("empty a")
	}

	m = len(a[0])
	if m < 1 {
		panic("empty a[0]")
	}
	for i := range a {
		if len(a[i]) != m {
			panic(fmt.Errorf("length of a[%d] != %d", i, m))
		}
	}

	return
}

func main() {
	// http://algolist.manual.ru/olimp/poi_prb.php
	a := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
	}

	n, m := getSize(a)
	x := 15
	i, j := n-1, m-1

	for ; i >= 0 && a[i][0] > x; i-- {
	}
	for ; j >= 0 && a[i][j] > x; j-- {
	}

	if a[i][j] == x {
		println(i, j)
	} else {
		println("НЕТ")
	}
}
