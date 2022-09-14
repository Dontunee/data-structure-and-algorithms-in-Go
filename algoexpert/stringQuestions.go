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
