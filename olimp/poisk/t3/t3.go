package main

import "git.lan.316.su/go/algo/shared"

type index struct {
	i, j int
}

func findElements(a [][]int) []index {
	idxs := make([]index, 0)

	n, m := shared.GetSize2D(a)
	for i := 0; i < n && a[i][0] >= i; i++ {
		for j := 0; j < m && a[i][j] >= i+j; j++ {
			if a[i][j] == i+j {
				idxs = append(idxs, index{i, j})
			}
		}
	}

	return idxs
}

func main() {
	// http://algolist.manual.ru/olimp/poi_prb.php (задача 3-4)
	a := [][]int{
		{10, 9, 2, 0, -5},
		{9, 8, 7, 6, 5},
	}
	idxs := findElements(a)

	if len(idxs) == 0 {
		println("НЕТ")
	} else {
		for _, idx := range idxs {
			println(idx.i, idx.j)
		}
	}
}
