func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left := 0
	maxCount := 0
	longest := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		count[c]++
		maxCount = max(maxCount, count[c])
		windowLen := right -left + 1
		if windowLen-maxCount > k{
			count[s[left]]--
			left++
		}
		longest = max(longest, right-left+1)
	}
	return longest
}
