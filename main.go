package main

import "fmt"

func main() {
	x := 5
	y := 4

	fmt.Printf("%d + %d = %d\n", x, y, Add(x, y))
}

// Add sums one or more values
func Add(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
