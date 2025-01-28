func maxArea(height []int) int {
	max_water := 0
	l := 0
	r := len(height) - 1
	min_value := 0

	for l < r {
		if height[l] < height[r] {
			min_value = height[l]
		} else {
			min_value = height[r]
		}
		calculation := min_value * (r - l)
		if calculation > max_water {
			max_water = calculation
		}
		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}

	}
	return max_water

}