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

func NonDPFactorial(n int) int {
	if n < 2 {
		return n
	}
	return NonDPFactorial(n-1) + NonDPFactorial(n-2)
}

func MemoizationFactorial(n int) int {
	memo := make([]int, n+1)
	factorial := calculateMemoizationRecursive(memo, n)
	return factorial
}

func calculateMemoizationRecursive(memo []int, n int) int {
	if n < 2 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = calculateMemoizationRecursive(memo, n-1) + calculateMemoizationRecursive(memo, n-2)
	return memo[n]
}

func TabulationFibonacci(n int) int {
	if n == 0 {
		return 0
	}
	fibArray := make([]int, n+1)
	fibArray[0] = 0
	fibArray[1] = 1

	for i := 2; i <= n; i++ {
		fibArray[i] = fibArray[i-1] + fibArray[i-2]
	}
	return fibArray[n]
}
