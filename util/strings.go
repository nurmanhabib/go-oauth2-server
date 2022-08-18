package util

import "math/rand"

// SliceContains is a function to make sure a string is in the slice of string.
func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// SliceSubtract is the subtraction of Slice A to Slice B.
func SliceSubtract(a, b []string) []string {
	var r []string

	for _, v := range a {
		if !SliceContains(b, v) {
			r = append(r, v)
		}
	}

	return r
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomAlpha is function to generate random alphabet.
func RandomAlpha(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
