package leetcode424

import "testing"

func TestCase1(t *testing.T) {
	result := characterReplacement("ABAB", 2)
	if result != 4 {
		t.Errorf("got %d, want %d", result, 4)
	}
}

func TestCase2(t *testing.T) {
	result := characterReplacement("AABABBA", 1)
	if result != 4 {
		t.Errorf("got %d, want %d", result, 4)
	}
}

func TestCase3(t *testing.T) {
	result := characterReplacement("ABCD", 2)
	if result != 3 {
		t.Errorf("got %d, want %d", result, 3)
	}
}
