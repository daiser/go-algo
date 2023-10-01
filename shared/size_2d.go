package shared

import "fmt"

func GetSize2D[T any](a [][]T) (n, m int) {
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
