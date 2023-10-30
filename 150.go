// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3

// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// Constraints:

//     n == nums.length
//     1 <= n <= 5 * 104
//     -109 <= nums[i] <= 109

// Follow-up: Could you solve the problem in linear time and in O(1) space?

package ll150

import (
// "math"
)

func MajorityElement(nums []int) (int, error) {
	//var es [150]int
	//mid := math.Floor(len(nums) / 2)
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++ // = m[nums[i]] + 1
	}
	maxVind := 0
	for k, v := range m {
		if v > m[maxVind] {
			maxVind = k
		}
	}
	return maxVind, nil
}
