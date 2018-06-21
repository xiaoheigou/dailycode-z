package longeststr

func lengthOfNonReatSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}
