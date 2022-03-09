package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	FindUnPairedElement([]int{9, 3, 9, 3, 9, 7, 9})
}

func PrintTriangle(n int) {
	for i := 1; i <= n; i++ {
		for value := 0; value < i; value++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func PrintDaysOfTheWeek() {
	days := map[string]string{
		"mon":  "monday",
		"tue":  "tuesday",
		"wed":  "wednesday",
		"thur": "thursday",
		"fri":  "friday",
		"sat":  "saturday",
		"sun":  "sunday",
	}

	for key, value := range days {
		fmt.Println(key + " " + " stands for " + value)
	}
}

func MaximumGap(N int) int {
	maximumGap := math.MinInt
	currentCount := 0
	inputBinary := strconv.FormatInt(int64(N), 2)
	runeSample := []rune(inputBinary)
	for i := 0; i < len(runeSample); i++ {
		if runeSample[i] == '1' {
			i++
			for ; i < len(runeSample); i++ {
				if inputBinary[i] == '0' {
					currentCount++
				} else if inputBinary[i] == '1' {
					maximumGap = int(math.Max(float64(currentCount), float64(maximumGap)))
					currentCount = 0
				}
			}
		}
	}
	return maximumGap
}

func ReverseArray(input []int) []int {
	length := len(input)
	for i := 0; i < length/2; i++ {
		k := length - i - 1
		input[i], input[k] = input[k], input[i]
	}
	return input
}

func FindUnPairedElement(A []int) int {
	var arrayValue int
	dictionary := make(map[int]int, 0)
	for _, value := range A {
		dictionary[value]++
	}

	for _, arrayValue = range A {
		if figure, ok := dictionary[arrayValue]; ok {
			if figure == 1 {
				return arrayValue
			}
		}
	}

	return arrayValue
}

func RotateNTimes(A []int, K int) []int {
	// write your code in Go 1.4
	length := len(A)
	K = K % length
	reverse(A, 0, length-1)
	reverse(A, 0, K-1)
	reverse(A, K, length-1)
	return A
}

func reverse(input []int, start int, end int) []int {
	for start < end {
		temp := input[start]
		input[start] = input[end]
		input[end] = temp
		start++
		end--
	}
	return input
}
