package question_75

func sortColors(nums []int) {
	/*
	   * Input:
	       - nums: []int -> list of numbers containing 0, 1, 2
	   * Process:
	       - Sort so that objects with the same color are adjacent
	   * Output:
	       - return: none -> Sort in-place
	   * Constraints:
	       - len(nums) <= 300
	       => So most likely will work with any solution
	       - One pass, constant extra space
	       =>  Runtime: O(N)
	           Space: O(1)
	   * Approach:
	       - For this problem, we can see that at the final solution we can see that 0 and 2 are on the two side
	       of the array
	       - With that being said, we can set 2 pointers, one 1 zero and 1 for two at the two ends of the array.
	       - We do not need to set pointer for one since we are not sure where one sequence will start.
	       - Every time that we see 0 or 2, we will swap the number with corresponding index of the 2 pointers.
	       - With that being said, we will eventually swap all the array until reaching the end index of the
	       array.
	   * Result:
	       - Runtime: O(N)
	       - Space: O(1)
	*/

	zeroIndex, twoIndex := 0, len(nums)-1
	index := 0

	for index < len(nums) {
		num := nums[index]
		for zeroIndex < len(nums) && nums[zeroIndex] == 0 {
			zeroIndex++
		}

		for twoIndex >= 0 && nums[twoIndex] == 2 {
			twoIndex--
		}

		if num == 2 && index < twoIndex {
			nums[index], nums[twoIndex] = nums[twoIndex], nums[index]
			continue
		} else if num == 0 && index > zeroIndex {
			nums[index], nums[zeroIndex] = nums[zeroIndex], nums[index]
			continue
		}

		index++
	}

}
