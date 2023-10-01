package main

import "git.lan.316.su/go/algo/shared"

func main() {
	a := [][]int{
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 0, 1}, // the line
		{0, 1, 0, 0},
	}

	n, m := shared.GetSize2D(a)
	if n != m {
		panic("square table required (NxN)")
	}

	row := 0
	for col := 1; col < n; col++ {
		if a[row][col] == 0 {
			row = col
		}
	}

	found := true
	// checking row
	for col := 0; col < n; col++ {
		found = found && (a[row][col] == 1 || row == col)
	}
	// checking column
	for j := 0; j < n; j++ {
		found = found && a[j][row] == 0
	}

	if found {
		println(row)
	} else {
		println("НЕТ")
	}
}
