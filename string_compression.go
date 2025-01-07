package main

import "fmt"

func StringCompression(chars []byte) int {
	i, k, n := 0, 0, len(chars)

	for i < n {
		var j = i + 1

		for j < n && chars[j] == chars[i] {
			j += 1

		}
		chars[k] = chars[i]
		k += 1
		if j-i > 1 {
			// print(j - i)
			count := fmt.Sprintf("%d", j-i)

			for _, c := range count {
				chars[k] = byte(c)
				k += 1

			}
		}
		i = j
	}
	return k
}
