package leetcode3

import "testing"

func TestCase1(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")
	if result != 3 {
		t.Errorf("got %d, want %d", result, 3)
	}
}

// abcadcbb
func TestCase2(t *testing.T) {
	result := lengthOfLongestSubstring(" ")
	if result != 1 {
		t.Errorf("got %d, want %d", result, 1)
	}
}

func TestCase3(t *testing.T) {
	result := lengthOfLongestSubstring("bbbbb")
	if result != 1 {
		t.Errorf("got %d, want %d", result, 1)
	}
}

func TestCase4(t *testing.T) {
	result := lengthOfLongestSubstring("pwwkew")
	if result != 3 {
		t.Errorf("got %d, want %d", result, 3)
	}
}

func TestCase5(t *testing.T) {
	result := lengthOfLongestSubstring("abba")
	if result != 2 {
		t.Errorf("got %d, want %d", result, 2)
	}
}
