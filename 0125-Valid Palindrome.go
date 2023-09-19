package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	l := 0          // Left pointer
	r := len(s) - 1 //Right pointer
	for l < r {

		left := unicode.ToLower(rune(s[l]))
		right := unicode.ToLower(rune(s[r]))

		//The following two if statements are responsible for skipping past non alphanumeric characters

		if !isAlphanumeric(left) {
			l++
			continue
		}

		if !isAlphanumeric(right) {
			r--
			continue
		}

		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}
