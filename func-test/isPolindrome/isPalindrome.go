package isPolindrome

func IsPalindrome(text string) bool {
	if len(text) <= 1 {
		return true
	}
	if len(text) == 2 && text[0] == text[1] {
		return true
	}
	if text[0] != text[len(text)-1] {
		return false
	}
	return IsPalindrome(text[1 : len(text)-1])
}
