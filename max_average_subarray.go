package main

func MaximunAverageSubarray(nums []int, k int) float64 {

	sum_of_window := 0

	for i := range len(nums[:k]) {
		sum_of_window += nums[i]
	}

	max_sum := sum_of_window

	for i := k; i < len(nums); i++ {
		sum_of_window += nums[i]
		sum_of_window -= nums[i-k]
		if sum_of_window > max_sum {
			max_sum = sum_of_window
		}
	}

	return float64(max_sum) / float64(k)

}
