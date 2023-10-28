package two_sum

/*
SPACE - TIME Complexity

Time - Complexity: O(n)
	buildMap() => O(1)
	Iterating over nums => O(n)

Space - Complexity: O(1)
*/

/*
 - Return the INDICESE which add up to target
 - May not use the same element twice
 - Exactly one solution
*/

func twoSum(nums []int, target int) []int {
	/*
		Build map
			Key => int from nums
			Val => index

		For index, num in nums
			search for 'target - num' in hashmap

			if exists
				slice of index and hashmap value
	*/

	hash := buildMap(nums, target)

	for index, num := range nums {
		t := target - num
		value, exists := hash[t]

		if exists && index != value {
			return []int{index, value}
		}
	}

	return nil
}

func buildMap(nums []int, target int) map[int]int {
	hash := make(map[int]int)

	for index, num := range nums {
		hash[num] = index
	}

	return hash
}
