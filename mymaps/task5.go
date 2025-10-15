package mymaps

func cachedFibonacci() func(int) int {

	cache := make(map[int]int)

	var fib func(n int) int
	fib = func(n int) int {
		if cache[n] != 0 {
			return cache[n]
		}

		if n == 0 {
			cache[n] = 0
			return 0
		}

		if n == 1 {
			cache[n] = 1
			return 1
		}

		sum := fib(n-1) + fib(n-2)
		cache[n] = sum
		return sum
	}
	return fib
}
