package ui

import "strings"

func isOneLetterAlpha(s string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range alpha {
		if string(c) == s {
			return true
		}
	}
	return false
}

func toUpperRune(s string) rune {
	return rune(strings.ToUpper(s)[0])
}

func lastRune(s string) rune {
	if s == "" {
		return '0'
	}
	return rune(s[len(s)-1])
}

// require: 'A' <= r <= 'Z'
func nextAlpha(r rune) rune {
	if r == 'Z' {
		return 'A'
	}
	return r + 1
}

// require: 'A' <= r <= 'Z'
func prevAlpha(r rune) rune {
	if r == 'A' {
		return 'Z'
	}
	return r - 1
}
