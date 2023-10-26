package contains_duplicates

func containsDuplicates(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}
	return false
}
