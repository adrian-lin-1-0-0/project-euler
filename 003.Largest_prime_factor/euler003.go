package euler003

import "math"

var primeCache = map[int64]bool{2: true, 3: true}

func isPrime(n int64) bool {
	if n&1 == 0 || n < 2 {
		return false
	}
	if primeCache[n] {
		return true
	}
	sqrtN := int64(math.Sqrt(float64(n))) + 1

	for i := int64(3); i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	primeCache[n] = true
	return true
}

func LargestPrimeFactor(n int64) int64 {
	sqrtN := int64(math.Sqrt(float64(n))) + 1
	max := int64(1)
	for i := int64(1); i < sqrtN; i++ {
		if n%i == 0 {
			d := n / i
			if isPrime(d) {
				return d
			}

			if isPrime(i) {
				max = i
			}
		}
	}
	return max
}
