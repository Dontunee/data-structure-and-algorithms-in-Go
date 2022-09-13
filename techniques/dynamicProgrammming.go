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

func BruteForceKnapsack(weights, profit []int, capacity int) int {
	//store maximum capacity
	//loop through weights
	//for each weight
	//check if current weight picked is higher than maxCap
	//if it is update max cap
	// then continue to add to each other weight
	// if weight is less than or equal to cap
	// add profits together
	// compare addition to existing max cap
	// update added figure
	//return max cap

	var maxCap int
	// abcd
	// 1,2,3,5
	// 1,6,10,16
	// c-> 7
	for weightKey, weightValue := range weights {
		if weightValue <= capacity {
			if profit[weightKey] > maxCap {
				maxCap = profit[weightKey]
			}

			for i := 0; i < len(weights); i++ {
				if i != weightKey {
					currentWeight := weights[weightKey] + weights[i]
					if currentWeight <= capacity {
						currentProfit := profit[weightKey] + profit[i]
						if currentProfit > maxCap {
							maxCap = currentProfit
						}
					}
				}
			}
		}
	}
	return maxCap
}

func RecursionKnapsack(weights, profits []int, capacity int) int {
	// abcd
	// 1,2,3,5
	// 1,6,10,16
	// c-> 7
	return recursionKnapsack(weights, profits, capacity, 0)
}

func recursionKnapsack(weights []int, profits []int, capacity, currentIndex int) int {
	//base checks
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}
	profitOne := 0
	// recursive call after choosing the element at the currentIndex
	// if the weight of the element at currentIndex exceeds the capacity, we shouldn't process this
	if weights[currentIndex] <= capacity {
		profitOne = recursionKnapsack(weights, profits, capacity-weights[currentIndex], currentIndex+1)
	}

	// recursive call after excluding the element at the currentIndex
	profitTwo := recursionKnapsack(weights, profits, capacity, currentIndex)

	return max(profitOne, profitTwo)

}

func max(numberOne, numberTwo int) int {
	if numberOne > numberTwo {
		return numberOne
	}
	return numberTwo
}

func Knapsack2dProblem(items [][]int, capacity int) []interface{} {
	return recursionSolution(items, capacity, 0, 0, []int{}, len(items))
}

func recursionSolution(items [][]int, capacity, i, currentValue int, indices []int, n int) []interface{} {
	if capacity <= 0 || n <= 0 {
		return []interface{}{currentValue, []int{}}
	}

	value := items[i][0]
	weight := items[i][1]

	exclude := []interface{}{0, []int{}}
	if capacity > weight {
		exclude = recursionSolution(items, capacity, i-1, currentValue, indices, n-1)
	}

	include := recursionSolution(items, capacity-weight, i-1, currentValue+value, indices, n-1)

	if include[0].(int) >= exclude[0].(int) {
		return []interface{}{include[0], include[1]}
	}
	return []interface{}{exclude[0], exclude[1]}
}
