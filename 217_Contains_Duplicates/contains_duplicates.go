package containsDuplicates

/*
SPACE - TIME Complexity

Time - Complexity: O(n)
Space - Complexity: 0(1)
*/

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}
	return false
}
