package main

import "fmt"

func MoveZeros(nums []int) []int {
	array_first := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[array_first], nums[i] = nums[i], nums[array_first]
			array_first += 1
		}
	}
	fmt.Println(nums)
	return nums
}
