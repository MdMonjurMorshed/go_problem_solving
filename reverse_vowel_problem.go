package main

import "strings"

func ReverseVowelProblemLeetcode345(s string) string {
	arr_str := []rune(s)
	vowel := "AEIOUaeiou"
	left := 0
	right := len(arr_str) - 1

	for i := 0; i < len(arr_str); i++ {
		if left < right {
			if !strings.Contains(vowel, string(arr_str[left])) {
				left += 1
			}
			if !strings.Contains(vowel, string(arr_str[right])) {
				right -= 1
			}
			if strings.Contains(vowel, string(arr_str[left])) && strings.Contains(vowel, string(arr_str[right])) {
				arr_str[left], arr_str[right] = arr_str[right], arr_str[left]
				left += 1
				right -= 1
			}
		}
	}

	return string(arr_str)

}


