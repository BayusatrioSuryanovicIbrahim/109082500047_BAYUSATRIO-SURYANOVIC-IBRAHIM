package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n *int) {
	var val int
	*n = 0

	for *n < 2022 {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}

		if exist(*T, *n, val) {
			break
		}

		T[*n] = val
		*n++
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0

	for i := 0; i < n; i++ {
		val := T1[i]
		if exist(T2, m, val) {
			T3[*h] = val
			*h++
		}
	}
}

func main() {
	var T1, T2, T3 set
	var n, m, h int

	inputSet(&T1, &n)
	inputSet(&T2, &m)

	findIntersection(T1, T2, n, m, &T3, &h)

	for i := 0; i < h; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T3[i])
	}
	fmt.Println()
}
