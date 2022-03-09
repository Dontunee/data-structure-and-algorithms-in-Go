package techniques

func NSum(n int) int {
	cache := make(map[int]int)
	if cache[n] > 0 {
		return cache[n]
	}
	result := -1
	if n < 2 {
		result = 1
	} else {
		result = NSum(n-1) + n
		return result
	}
	cache[n] = result
	return result
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	result := Factorial(n-1) * n
	return result
}
