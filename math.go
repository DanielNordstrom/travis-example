package main

// Add sums one or more values
func Add(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

// Sub subtracts numbers from the first number
func Sub(nums ...int) (total int) {
	total = nums[0]

	for _, num := range nums[1:] {
		total -= num
	}
	return
}
