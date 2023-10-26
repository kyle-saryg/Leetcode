package valid_anagram

/*
Same number of characters
map{char: int (number of occurences)}
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := createMap(s)
	tMap := createMap(t)

	for sKey, sValue := range sMap {
		tValue, tExists := tMap[sKey]
		if !tExists || tValue != sValue {
			return false
		}
	}
	return true
}

// Create hash map
func createMap(foo string) map[rune]int {
	dict := make(map[rune]int)

	for _, char := range foo {
		dict[char]++
	}

	return dict
}
