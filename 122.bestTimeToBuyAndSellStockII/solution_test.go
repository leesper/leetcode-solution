package leetcode122

import "testing"

func TestCase1(t *testing.T) {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if result != 7 {
		t.Errorf("got %d, want %d", result, 7)
	}
}

func TestCase2(t *testing.T) {
	result := maxProfit([]int{1, 2, 3, 4, 5})
	if result != 4 {
		t.Errorf("got %d, want %d", result, 4)
	}
}

func TestCase3(t *testing.T) {
	result := maxProfit([]int{7, 6, 4, 3, 1})
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}
