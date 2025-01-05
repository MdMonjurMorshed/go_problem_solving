package main

import "fmt"

// import "fmt"

func IncreasingTripletSubsecquence(nums []int) bool {

	if len(nums) < 3 {
		return false
	}

	var i = (1 << 31) - 1
	var j = (1 << 31) - 1
	fmt.Println(i)
	for num := 0; num < len(nums); num++ {
		if nums[num] <= i {
			fmt.Println(num)
			i = nums[num]
		} else if nums[num] <= j {
			fmt.Println(num)
			j = nums[num]
		} else {
			return true
		}
	}

	return false
}
