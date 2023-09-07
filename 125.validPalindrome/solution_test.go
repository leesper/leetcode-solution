package leetcode125

import "testing"

func TestCase1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	if !isPalindrome(s) {
		t.Fatalf("expected to be true")
	}
}

func TestCase2(t *testing.T) {
	s := "race a car"
	if isPalindrome(s) {
		t.Fatalf("expected to be false")
	}
}

func TestCase3(t *testing.T) {
	s := " "
	if !isPalindrome(s) {
		t.Fatalf("expected to be true")
	}
}
