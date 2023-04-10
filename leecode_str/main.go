package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func countSubstrings(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += countPalindromes(s, i, i)
		ans += countPalindromes(s, i, i+1)
	}
	return ans
}

func countPalindromes(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		count++
	}
	return count
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphaNumeric(s[l]) {
			l++
			continue
		}
		if !isAlphaNumeric(s[r]) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func main() {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(res)
}
