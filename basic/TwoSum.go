package basic

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example:
//
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]
func twoSum(nums []int, target int) []int {
	for indexX := len(nums) - 1; indexX > 0; indexX-- {
		x := nums[indexX]
		for indexY := 0; indexY < indexX; indexY++ {
			if (x + nums[indexY]) == target {
				return []int{indexY, indexX}
			}
		}
	}
	return nil
}

// run result time: 40ms memory 3M
