package palindrome

import "strings"

func Palindrome(s string) bool {
	var bol bool
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if s[i] == s[len(s)-1-i] {
			bol = true
		} else {
			bol = false
		}
	}
	return bol
}
