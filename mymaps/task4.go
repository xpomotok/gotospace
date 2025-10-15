package mymaps

func invertMap(original map[string]int) map[int][]string {

	inv := make(map[int][]string)

	for k, v := range original {
		inv[v] = append(inv[v], k)
	}

	return inv
}
