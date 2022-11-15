package isPolindrome

import (
	"testing"
)

func TestIsPalindrome_WithLengthLessThanOrEqualOne(t *testing.T) {
	expectation := true
	actual := IsPalindrome("a")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsPalindrome_WithLengthExactlyTwoAndSameCharacter(t *testing.T) {
	expectation := true
	actual := IsPalindrome("aa")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsPalindrome_WithFirstCharAndLastCharAreDifferent(t *testing.T) {
	expectation := false
	actual := IsPalindrome("abc")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsPalindrome_WithPalindromeText(t *testing.T) {
	expectation := true
	actual := IsPalindrome("kayak")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsPalindrome_WithNoPalindromeText(t *testing.T) {
	expectation := false
	actual := IsPalindrome("palindrome")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
