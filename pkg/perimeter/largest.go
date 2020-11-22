package perimeter

import (
	"sort"
)

// Given an array A of positive lengths, return the largest perimeter
// of a triangle with non-zero area, formed from 3 of these lengths.
// If it is impossible to form any triangle of non-zero area, return 0.

func LargestPerimeter(A []int) int {
	// Check valid triangle
	if len(A) < 3 {
		return 0
	}
	// Sort slice
	sort.Ints(A)
	// Iterate over lengths
	// check if non-zero
	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
