package basic

// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
//
// Input: 123
// Output: 321

// Example 2:
//
// Input: -123
// Output: -321

// Example 3:
//
// Input: 120
// Output: 21

// the function's result is failed
// the function has not border judgment of int32 min
func reverse1(x int) int {
	var result int
	for x != 0 {
		result = (result * 10) + (x % 10)
		x = x / 10
	}
	return result
}

// optimize reverse1
// the function's result is pass
// result infos: 1032 / 1032 cases, running time 4 ms
func reverse(x int) int {
	result := 0
	const Int32Maximum = 1<<31 - 1
	const Int32Minimum = -1 << 31
	for x != 0 {
		result = (result * 10) + (x % 10)
		if result <= Int32Minimum || result >= Int32Maximum {
			return 0
		}
		x = x / 10
	}
	return result
}
