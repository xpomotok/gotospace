package mymaps

import "slices"

func findIntersection(slice1, slice2 []int) []int {

	dict := make(map[int]int)

	for _, v := range slice1 {

		dict[v]++
	}

	for _, v := range slice2 {
		dict[v]++
	}

	slice3 := make([]int, 0)

	for k, v := range dict {
		if v >= 2 {
			slice3 := append(slice3, k)

		}
	}

	slices.Sort(slice3)

	return slice3
}
