package isPolilndromeLoop

import (
	"testing"
)

func TestIsPalindromeLoopWithPalindromeText(t *testing.T) {
	expectation := true
	actual := IsPalindrineLoop("kasurnababanrusak")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsPalindromeLoopWithNonPalindromeText(t *testing.T) {
	expectation := false
	actual := IsPalindrineLoop("palindrome")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
