package algoexpert

import (
	"strconv"
	"strings"
)

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

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	solution := caesarCipherEncryptor(str, key)
	return solution
}

func caesarCipherEncryptor(str string, key int) string {
	const alphabets = "abcdefghijklmnopqrstuvwxyz"
	alphaArray := []rune(alphabets)

	sAR := []rune(str)
	var builder strings.Builder
	for _, v := range sAR {
		indexOf := indexOf(v, alphaArray)

		newIndex := (indexOf + key) % 26
		builder.WriteRune(alphaArray[newIndex])
	}
	return builder.String()
}

func indexOf(r rune, input []rune) int {
	for p, v := range input {
		if v == r {
			return p
		}
	}
	return 0
}

func RunLengthEncoding(str string) string {
	// Write your code here.
	return runLengthEncoding(str)
}

func runLengthEncoding(str string) string {

	counter := 1
	var sb strings.Builder
	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		prevChar := str[i-1]

		if currentChar != prevChar || counter == 9 {
			//add count to characters
			//add previousChar to characters
			sb.WriteString(strconv.Itoa(counter))
			sb.WriteRune(rune(prevChar))
			counter = 0
		}
		counter++
	}

	//add count to characters
	//add last element to characters
	sb.WriteString(strconv.Itoa(counter))
	sb.WriteRune(rune(str[len(str)-1]))

	//conv characters to string
	return sb.String()
}