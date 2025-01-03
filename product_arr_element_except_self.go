package main

import "fmt"

func ProductOfArrayElementExceptSlef(nums []int) []int {

	ans := make([]int, len(nums))
	left := 1
	right := 1

	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		ans[i] = left
		left *= nums[i]
	}
	fmt.Println(ans)
	for j := len(nums) - 1; j >= 0; j-- {
		ans[j] *= right
		right *= nums[j]
	}
	fmt.Println(ans)
	return ans

}

func ProductOfArrayElementExceptSlefAnother(nums []int) []int {
	var ans []int
	mul := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans = append(ans, 1)
			continue

		}

		ans = append(ans, ans[i-1]*nums[i-1])

	}

	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			ans[j] = ans[j] * mul
			continue

		}
		mul = mul * nums[j+1]
		ans[j] = ans[j] * mul

	}
	return ans

}
