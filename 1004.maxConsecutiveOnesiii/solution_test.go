package leetcode1004

import "testing"

func TestCase1(t *testing.T) {
	result := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	if result != 6 {
		t.Errorf("got %d, want %d", result, 6)
	}
}

func TestCase2(t *testing.T) {
	result := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	if result != 10 {
		t.Errorf("got %d, want %d", result, 10)
	}
}

func TestCase3(t *testing.T) {
	result := longestOnes([]int{0, 0, 0, 0}, 0)
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}

func TestCase4(t *testing.T) {
	result := longestOnes([]int{0, 0, 0, 1}, 4)
	if result != 4 {
		t.Errorf("got %d, want %d", result, 4)
	}
}

func TestCase5(t *testing.T) {
	result := longestOnes([]int{0, 1, 1}, 0)
	if result != 2 {
		t.Errorf("got %d, want %d", result, 2)
	}
}
