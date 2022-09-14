package algoexpert

import "strings"

func IsPalindrome(str string) bool {
	// Write your code here.
	return buildFromBack(str)
}

func buildFromBack(str string) bool {
	strArray := []rune(str)
	if len(strArray) <= 0 {
		return false
	}
	var builder strings.Builder
	for i := len(strArray) - 1; i >= 0; i-- {
		builder.WriteString(string(strArray[i]))
	}

	if str == builder.String() {
		return true
	}

	return false
}

// O(n) time complexity and O(n) space complexity
func recursive(str string, i, j int) bool {
	//base case
	if i >= j {
		return true
	}

	//General case
	if str[i] != str[j] {
		return false
	}

	return recursive(str, i+1, j-1)
}

// O(n) time complexity and O(1) space complexity
func pointer(str string) bool {
	// Write your code here.
	j := len(str) - 1
	for i := 0; i < len(str); i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}

	return true
}
