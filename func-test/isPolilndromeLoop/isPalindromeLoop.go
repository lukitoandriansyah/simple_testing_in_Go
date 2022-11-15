package isPolilndromeLoop

func IsPalindrineLoop(text string) bool {
	n := len(text) / 2
	for i := 0; i < n; i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}
