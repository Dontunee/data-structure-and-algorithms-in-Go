package main

import (
	"fmt"
	"github.com/dontunee/algorithmInGo/blind75"
	"math"
	"sort"
	"strconv"
)

func main() {
	blind75.MaxProfit([]int{2, 1, 2, 1, 0, 1, 2})
}

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i++ {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	result := make([]int, length+1)
	result[0] = 1
	return result
}
func SortedSquares(nums []int) []int {
	length := len(nums)
	firstPointer := 0
	secondPointer := length - 1
	output := make([]int, length)
	counter := length - 1

	for firstPointer <= secondPointer {
		firstSquare := nums[firstPointer] * nums[firstPointer]
		secondSquare := nums[secondPointer] * nums[secondPointer]

		if firstSquare > secondSquare {
			output[counter] = firstSquare
			firstPointer++
		} else {
			output[counter] = secondSquare
			secondPointer--
		}
		counter--
	}

	return output

}

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for index, value := range nums {
		if index > 0 && value == nums[index-1] {
			continue
		}

		leftPointer, rightPointer := index+1, len(nums)-1

		for leftPointer < rightPointer {
			threeSum := value + nums[leftPointer] + nums[rightPointer]

			if threeSum > 0 {
				rightPointer--
			} else if threeSum < 0 {
				leftPointer++
			} else {
				newArray := []int{value, nums[leftPointer], nums[rightPointer]}
				result = append(result, newArray)
				leftPointer++

				for nums[leftPointer] == nums[leftPointer-1] && leftPointer < rightPointer {
					leftPointer += 1
				}

			}

		}
	}

	return result
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

func ContainsDuplicate(nums []int) bool {
	length := len(nums)
	response := false
	if length < 1 || length > int(math.Pow(float64(10), float64(5))) {
		return false
	}

	dictionary := make(map[int]int, 0)
	for _, value := range nums {
		dictionaryValue := dictionary[value]
		if dictionaryValue > 0 {
			response = true
			return response
		}
		dictionary[value]++
	}
	return response
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
