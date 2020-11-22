// Package runningsum solves https://leetcode.com/problems/running-sum-of-1d-array/
package runningsum

// RunningSum https://leetcode.com/problems/running-sum-of-1d-array/
func RunningSum(nums []int) []int {
	var (
		ints []int
	)

	if len(nums) < 1000 {
		ints = append(ints, nums[0])
		for i := 1; i < len(nums); i++ {
			ints = append(ints, nums[i]+ints[i-1])
		}
	}

	return ints
}
