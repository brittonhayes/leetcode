package maxones

func FindMaxConsecutiveOnes(nums []int) int {

	// Set maximum consecutive
	max := 0
	current := 0

	// Range through array
	for _, n := range nums {
		if n != 0 {
			// Increase current max
			// if no zero
			current++
		} else {
			// Reset current max
			// otherwise
			current = 0
			continue
		}

		// Check for max consecutive
		// ones and return
		if max < current {
			max = current
			continue
		}
	}

	return max
}
