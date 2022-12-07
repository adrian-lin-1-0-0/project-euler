package euler002

var memCache = map[int64]int64{}

func FibEven(n int64) int64 {
	sum := int64(0)
	for i := int64(1); i <= n; i++ {
		r := fib(i)
		if r < n {
			if r%2 == 0 {
				sum += r
			}
		} else {
			break
		}

	}
	return sum
}

func fib(n int64) int64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	ans, ok := memCache[n]
	if ok {
		return ans
	}

	memCache[n] = fib(n-1) + fib(n-2)
	return memCache[n]
}
